package domains

import "time"

type Products struct {
	ProductId         int              `gorm:"primary_key;column:product_id;auto_increment"`
	ListingUserId     int              `gorm:"column:listing_user_id"`
	CategoryId        int              `gorm:"column:category_id"`
	Title             string           `gorm:"column:title"`
	ProductSlug       string           `gorm:"column:product_slug"`
	Description       string           `gorm:"column:description"`
	Price             float64          `gorm:"column:price"`
	Condition         ProductCondition `gorm:"column:condition"`
	Status            ProductStatus    `gorm:"column:status"`
	ViewCount         int              `gorm:"column:view_count"`
	CreatedAt         time.Time        `gorm:"column:created_at"`
	UpdatedAt         time.Time        `gorm:"column:update_at"`
	ProductCategories []Categories     `gorm:"foreignKey:category_id;references:category_id"`
	ProductImages     []ProductImages  `gorm:"foreignKey:product_id;references:product_id"`
	ProductSpecs      []ProductSpecs   `gorm:"foreignKey:product_id;references:product_id"`
}

type ProductCondition string

const (
	New        ProductCondition = "New"
	LikeNew    ProductCondition = "Like New"
	Good       ProductCondition = "Good"
	GoodQuite  ProductCondition = "Good Quite"
	NeedRepair ProductCondition = "NeedRepair"
)

type ProductStatus string

const (
	Available ProductStatus = "Available"
	SoldOut   ProductStatus = "Sold out"
)

func (Products) TableName() string {
	return "products"
}
