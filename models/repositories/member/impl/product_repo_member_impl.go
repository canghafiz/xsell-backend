package impl

import (
	"be/helpers"
	"be/models/domains"
	"errors"

	"gorm.io/gorm"
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

	updates := make(map[string]interface{})

	if product.Title != "" {
		updates["title"] = product.Title
		slug := product.ProductSlug
		if slug == "" {
			slug = helpers.GenerateSlug(product.Title)
		}
		updates["product_slug"] = slug
	}
	if product.Description != "" {
		updates["description"] = product.Description
	}
	if product.Price != 0 {
		updates["price"] = product.Price
	}
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

	if product.ProductImages != nil {
		if err := tx.Where("product_id = ?", product.ProductId).Delete(&domains.ProductImages{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if len(product.ProductImages) > 0 {
			for i := range product.ProductImages {
				product.ProductImages[i].ProductId = product.ProductId
				product.ProductImages[i].ImageId = 0 // reset ID agar auto-increment
			}
			if err := tx.Create(&product.ProductImages).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if product.ProductSpecs != nil {
		if err := tx.Where("product_id = ?", product.ProductId).Delete(&domains.ProductSpecs{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if len(product.ProductSpecs) > 0 {
			for i := range product.ProductSpecs {
				product.ProductSpecs[i].ProductId = product.ProductId
				product.ProductSpecs[i].ProductSpecId = 0
			}
			if err := tx.Create(&product.ProductSpecs).Error; err != nil {
				tx.Rollback()
				return err
			}
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
