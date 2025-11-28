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
	"sort"
	"strings"

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

func (serv *PageLayoutServMemberImpl) GetHomeLayouts(sectionLimits map[string]int) (resources.PageResource, error) {
	return serv.getPageLayout("home", sectionLimits)
}

func (serv *PageLayoutServMemberImpl) GetProductDetailLayouts(sectionLimits map[string]int) (resources.PageResource, error) {
	return serv.getPageLayout("product_detail", sectionLimits)
}

func (serv *PageLayoutServMemberImpl) getPageLayout(pageKey string, sectionLimits map[string]int) (resources.PageResource, error) {
	var cacheKeys []string
	for secKey, limit := range sectionLimits {
		cacheKeys = append(cacheKeys, fmt.Sprintf("%s:%d", secKey, limit))
	}
	sort.Strings(cacheKeys) // agar urutan konsisten
	cacheKeySuffix := strings.Join(cacheKeys, ",")
	cacheKey := fmt.Sprintf("page:%s:layout:%s", pageKey, cacheKeySuffix)

	// Get From Cache
	var results resources.PageResource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.(*resources.PageResource); ok {
			log.Printf("Cache hit for page layout: %s", cacheKey)
			return *result, nil
		}
	}

	// Call Repo
	layouts, errLayout := serv.PageLayoutRepo.GetPageLayout(serv.Db, pageKey)
	if errLayout != nil {
		log.Printf("[PageLayoutRepo.GetPageLayout] error: %v", errLayout)
		return resources.PageResource{
			PageKey: pageKey,
			Data:    nil,
		}, fmt.Errorf("failed to fetch page layout")
	}

	var sections []section.Resource

	for _, layout := range layouts {
		if layout.Section == nil {
			continue
		}

		sectionKey := layout.Section.SectionKey
		limitToUse := 10 // default

		if customLimit, exists := sectionLimits[sectionKey]; exists {
			limitToUse = customLimit
		}

		products, err := serv.getProductsByConfig(layout.Section.Config, limitToUse)
		if err != nil {
			log.Printf("Failed to get products for section %s: %v", sectionKey, err)
			continue
		}

		sections = append(sections, *section.ToResource(*layout.Section, products))
	}

	result := resources.PageResource{
		PageKey: pageKey,
		Data:    sections,
	}

	// Save to cache
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, result, 1440); errCache != nil {
		log.Printf("Failed to cache page layout for key %s: %v", cacheKey, errCache)
	}

	return result, nil
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
