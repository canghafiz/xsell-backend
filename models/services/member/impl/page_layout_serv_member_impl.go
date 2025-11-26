package impl

import (
	"be/helpers"
	"be/models/domains"
	"be/models/repositories"
	"be/models/resources"
	"be/models/resources/section"
	"be/models/services"
	"fmt"
	"log"

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
	return serv.getPageLayout("home", overrideLimit)
}

func (serv *PageLayoutServMemberImpl) GetProductDetailLayouts(overrideLimit int) (resources.PageResource, error) {
	return serv.getPageLayout("product_detail", overrideLimit)
}

func (serv *PageLayoutServMemberImpl) getPageLayout(pageKey string, overrideLimit int) (resources.PageResource, error) {
	cacheKey := fmt.Sprintf("page:%v:layout:limit:%d", pageKey, overrideLimit)

	// Get from cache
	var results resources.PageResource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.(*resources.PageResource); ok {
			log.Printf("Successfully get page layout form key %s.", cacheKey)
			return *result, nil
		}
	}

	// Get from repo
	layouts, errLayout := serv.PageLayoutRepo.GetPageLayout(serv.Db, pageKey)
	if errLayout != nil {
		log.Printf("[PageLayoutRepo.GetPageLayout] error: %v", errLayout)
		return resources.PageResource{
			PageKey: pageKey,
			Data:    nil,
		}, fmt.Errorf("user not found")
	}

	var sections []section.Resource
	for _, layout := range layouts {
		if layout.Section == nil {
			continue
		}

		products, err := serv.getProductsByConfig(layout.Section.Config, overrideLimit)
		if err != nil {
			continue
		}

		sections = append(sections, *section.ToResource(*layout.Section, products))
	}

	// Set to cache
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, layouts, 1440); errCache != nil {
		log.Printf("Failed to set cache for key %s: %v", cacheKey, errCache)
		return resources.PageResource{
			PageKey: pageKey,
			Data:    nil,
		}, fmt.Errorf("get page layout failed, please try again later")
	}

	return resources.PageResource{
		PageKey: pageKey,
		Data:    sections,
	}, nil
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
