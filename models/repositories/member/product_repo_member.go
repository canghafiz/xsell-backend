package member

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type ProductRepoMember interface {
	Create(db *gorm.DB, product domains.Products) error
	Update(db *gorm.DB, product domains.Products) error
	GetSingleBySlug(db *gorm.DB, slug string) (*domains.Products, error)
	Delete(db *gorm.DB, product domains.Products) error
}
