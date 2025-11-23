package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type CategoryRepoAdminImpl struct {
}

func NewCategoryRepoAdminImpl() *CategoryRepoAdminImpl {
	return &CategoryRepoAdminImpl{}
}

func (repo *CategoryRepoAdminImpl) Create(db *gorm.DB, category domains.Categories) error {
	return db.Create(&category).Error
}

func (repo *CategoryRepoAdminImpl) Update(db *gorm.DB, category domains.Categories) error {
	if category.CategoryId == 0 {
		return gorm.ErrInvalidValue
	}
	return db.Model(&domains.Categories{}).
		Where("category_id = ?", category.CategoryId).
		Select("category_name", "description", "icon").
		Updates(category).Error
}

func (repo *CategoryRepoAdminImpl) Delete(db *gorm.DB, category domains.Categories) error {
	if category.CategoryId == 0 {
		return gorm.ErrInvalidValue
	}
	return db.Where("category_id = ?", category.CategoryId).
		Delete(&domains.Categories{}).Error
}
