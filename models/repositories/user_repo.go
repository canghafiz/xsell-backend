package repositories

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(db *gorm.DB, user domains.Users) error
	Update(db *gorm.DB, user domains.Users) (*domains.Users, error)
	FindByEmail(db *gorm.DB, email string) (*domains.Users, error)
	CheckPasswordValid(db *gorm.DB, user domains.Users) (bool, error)
	ResetToken(db *gorm.DB, id int) error
}
