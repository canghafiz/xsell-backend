package domains

import (
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	CategoryId   int       `gorm:"primary_key;column:category_id;auto_increment"`
	CategoryName string    `gorm:"column:category_name"`
	CategorySlug string    `gorm:"column:category_slug"`
	Description  string    `gorm:"column:description"`
	Icon         string    `gorm:"column:icon"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:update_at"`
}

func (c *Categories) BeforeCreate(tx *gorm.DB) error {
	c.CategorySlug = generateSlug(c.CategoryName)
	return nil
}

func (c *Categories) BeforeUpdate(tx *gorm.DB) error {
	c.CategorySlug = generateSlug(c.CategoryName)
	return nil
}

func generateSlug(s string) string {
	slug := strings.ToLower(s)

	reg, _ := regexp.Compile("[^a-z0-9\\s]+")
	slug = reg.ReplaceAllString(slug, " ")

	slug = strings.Join(strings.Fields(slug), "_")

	return slug
}

func (Categories) TableName() string {
	return "categories"
}
