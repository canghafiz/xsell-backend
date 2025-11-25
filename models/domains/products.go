package domains

import (
	"strings"
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Products struct {
	ProductId         int              `gorm:"primary_key;column:product_id;auto_increment"`
	ListingUserId     int              `gorm:"column:listing_user_id"`
	CategoryId        int              `gorm:"column:category_id"`
	Title             string           `gorm:"column:title"`
	ProductSlug       string           `gorm:"column:product_slug;index:idx_product_slug,unique"`
	Description       string           `gorm:"column:description"`
	Price             float64          `gorm:"column:price"`
	Condition         ProductCondition `gorm:"column:condition"`
	Status            ProductStatus    `gorm:"column:status"`
	ViewCount         int              `gorm:"column:view_count"`
	CreatedAt         time.Time        `gorm:"column:created_at"`
	UpdatedAt         time.Time        `gorm:"column:updated_at"`
	ProductCategories []Categories     `gorm:"foreignKey:category_id;references:category_id"`
	ProductImages     []ProductImages  `gorm:"foreignKey:product_id;references:product_id"`
	ProductSpecs      []ProductSpecs   `gorm:"foreignKey:product_id;references:product_id"`
}

func (p *Products) BeforeCreate(tx *gorm.DB) error {
	p.generateSlug()
	return nil
}

func (p *Products) BeforeUpdate(tx *gorm.DB) error {
	p.generateSlug()
	return nil
}

func (p *Products) generateSlug() {
	if p.Title == "" {
		return
	}
	s := slug.Make(strings.TrimSpace(p.Title))
	p.ProductSlug = s
}

type ProductCondition string

const (
	New        ProductCondition = "New"
	LikeNew    ProductCondition = "Like New"
	Good       ProductCondition = "Good"
	GoodQuite  ProductCondition = "Good Quite"
	NeedRepair ProductCondition = "Needs Repair"
)

type ProductStatus string

const (
	Available ProductStatus = "Available"
	SoldOut   ProductStatus = "Sold out"
)

func (Products) TableName() string {
	return "products"
}
