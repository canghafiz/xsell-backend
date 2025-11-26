package impl

import (
	"be/models/domains"
	"be/models/repositories"
	"be/models/resources"
	"be/models/services"

	"gorm.io/gorm"
)

type PageLayoutServMemberImpl struct {
	Db             *gorm.DB
	RedisServ      services.RedisService
	PageLayoutRepo repositories.PageLayoutRepo
}

func NewPageLayoutServMemberImpl(db *gorm.DB, redisServ services.RedisService, pageLayoutRepo repositories.PageLayoutRepo) *PageLayoutServMemberImpl {
	return &PageLayoutServMemberImpl{Db: db, RedisServ: redisServ, PageLayoutRepo: pageLayoutRepo}
}

func (serv *PageLayoutServMemberImpl) GetHomeLayouts(overrideLimit int) (resources.PageResource, error) {
	//TODO implement me
	panic("implement me")
}

func (serv *PageLayoutServMemberImpl) GetProductDetailLayouts(overrideLimit int) (resources.PageResource, error) {
	//TODO implement me
	panic("implement me")
}

func (serv *PageLayoutServMemberImpl) getProductsByConfig(config domains.ContentConfig, overrideLimit int) ([]domains.Products, error) {
	query := serv.Db.Model(&domains.Products{}).
		Where("status = ?", domains.Available).
		Preload("ProductImages").
		Preload("ProductCategories")

	if config.Filters != nil {
		if condition, exists := config.Filters["condition"].(string); exists {
			query = query.Where("condition = ?", condition)
		}
		if categorySlug, exists := config.Filters["category_slug"].(string); exists {
			query = query.Joins("JOIN categories ON products.category_id = categories.category_id").
				Where("categories.category_slug = ?", categorySlug)
		}
		if maxPrice, exists := config.Filters["max_price"].(float64); exists {
			query = query.Where("price <= ?", maxPrice)
		}
	}

	// Handle predefined strategies
	if config.Strategy != "" {
		switch config.Strategy {
		case "newly_listed":
			query = query.Order("created_at DESC")
		case "trending":
			query = query.Order("view_count DESC")
		case "trusted_sellers":
			query = query.Order("created_at DESC")
		}
	}

	// Handle fixed product IDs
	if len(config.ProductIDs) > 0 {
		query = query.Where("product_id IN ?", config.ProductIDs)
	}

	// Determine final limit
	limit := 10
	if overrideLimit > 0 {
		limit = overrideLimit
	} else if config.Limit > 0 {
		limit = config.Limit
	}
	query = query.Limit(limit)

	var products []domains.Products
	err := query.Find(&products).Error
	return products, err
}
