package impl

import (
	"be/helpers"
	"be/models/domains"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepoMemberImpl struct{}

func NewProductRepoMemberImpl() *ProductRepoMemberImpl {
	return &ProductRepoMemberImpl{}
}

func (repo *ProductRepoMemberImpl) Create(db *gorm.DB, product domains.Products) error {
	return db.Create(&product).Error
}

func (repo *ProductRepoMemberImpl) Update(db *gorm.DB, product domains.Products) error {
	if product.ProductId == 0 {
		return gorm.ErrInvalidValue
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// ✅ Pastikan produk ada
	var exists bool
	if err := tx.Model(&domains.Products{}).
		Select("1").
		Where("product_id = ?", product.ProductId).
		Limit(1).
		Scan(&exists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !exists {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	// --- 1. Update kolom utama ---
	updates := make(map[string]interface{})

	// Karena Title, Description, dll bukan pointer, cek kosong aman
	if product.Title != "" {
		updates["title"] = product.Title
		updates["product_slug"] = helpers.GenerateSlug(product.Title)
	}
	if product.Description != "" {
		updates["description"] = product.Description
	}
	if product.Price != 0 {
		updates["price"] = product.Price
	}
	// Karena Condition/Status wajib (validate:"required"), pasti tidak kosong jika dikirim
	if product.Condition != "" {
		updates["condition"] = product.Condition
	}
	if product.Status != "" {
		updates["status"] = product.Status
	}
	if product.CategoryId != 0 {
		updates["category_id"] = product.CategoryId
	}

	if len(updates) > 0 {
		if err := tx.Model(&domains.Products{}).
			Where("product_id = ?", product.ProductId).
			Updates(updates).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// --- 2. Handle Images: Full Replace ---
	if product.ProductImages != nil {
		if err := tx.Where("product_id = ?", product.ProductId).Delete(&domains.ProductImages{}).Error; err != nil {
			tx.Rollback()
			return err
		}
		if len(product.ProductImages) > 0 {
			for i := range product.ProductImages {
				product.ProductImages[i].ProductId = product.ProductId
				product.ProductImages[i].ImageId = 0
			}
			if err := tx.Create(&product.ProductImages).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// --- 3. Handle Specs: Partial UPSERT ---
	if product.ProductSpecs != nil {
		for _, spec := range product.ProductSpecs {
			spec.ProductId = product.ProductId
			err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "product_id"}, {Name: "category_product_spec_id"}},
				DoUpdates: clause.Assignments(map[string]interface{}{"spec_value": spec.SpecValue}),
			}).Create(&spec).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// --- 4. Handle Location (Opsional) ---
	// Jika struct `product.Location` diisi, update atau insert
	// Contoh: jika Location memiliki data (misal latitude != 0), maka simpan
	if product.Location.ProductId != 0 || product.Location.Latitude != 0 || product.Location.Longitude != 0 {
		// Cek apakah sudah ada
		var existingLocation domains.Location
		err := tx.Where("product_id = ?", product.ProductId).First(&existingLocation).Error
		if errors.Is(gorm.ErrRecordNotFound, err) {
			// Insert baru
			product.Location.ProductId = product.ProductId
			if err := tx.Create(&product.Location).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else if err == nil {
			// Update existing
			updatesLoc := make(map[string]interface{})
			if product.Location.Latitude != 0 {
				updatesLoc["latitude"] = product.Location.Latitude
			}
			if product.Location.Longitude != 0 {
				updatesLoc["longitude"] = product.Location.Longitude
			}
			if len(updatesLoc) > 0 {
				if err := tx.Model(&existingLocation).Updates(updatesLoc).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		} else {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (repo *ProductRepoMemberImpl) GetSingleBySlug(db *gorm.DB, slug string) (*domains.Products, error) {
	if slug == "" {
		return nil, gorm.ErrInvalidValue
	}

	var product domains.Products
	err := db.Preload("ProductCategories").
		Preload("ProductImages").
		Preload("ProductSpecs.CategoryProductSpec").
		Preload("Location").
		Where("product_slug = ?", slug).
		First(&product).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepoMemberImpl) Delete(db *gorm.DB, product domains.Products) error {
	if product.ProductId == 0 {
		return gorm.ErrInvalidValue
	}

	var count int64
	db.Model(&domains.Products{}).Where("product_id = ?", product.ProductId).Count(&count)
	if count == 0 {
		return gorm.ErrRecordNotFound
	}

	return db.Delete(&domains.Products{}, product.ProductId).Error
}
