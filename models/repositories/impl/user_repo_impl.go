package impl

import (
	"be/helper"
	"be/models/domains"
	"fmt"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
}

func NewUserRepoImpl() *UserRepoImpl {
	return &UserRepoImpl{}
}

func (repo *UserRepoImpl) Create(db *gorm.DB, user domains.Users) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepoImpl) Update(db *gorm.DB, user domains.Users) (*domains.Users, error) {
	result := db.
		Model(&domains.Users{}).
		Where("id = ?", user.UserId).
		Updates(map[string]interface{}{
			"email":             user.Email,
			"role":              user.Role,
			"first_name":        user.FirstName,
			"last_name":         user.LastName,
			"photo_profile_url": user.PhotoProfileUrl,
		})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user with ID %d not found or no changes were made", user.UserId)
	}

	return &user, nil
}

func (repo *UserRepoImpl) FindByEmail(db *gorm.DB, email string) (*domains.Users, error) {
	var user domains.Users
	result := db.
		Model(&user).
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepoImpl) CheckPasswordValid(db *gorm.DB, user domains.Users) (bool, error) {
	result, err := repo.FindByEmail(db, user.Email)
	if err != nil {
		return false, err
	}

	return helper.CheckPassword(result.Password, user.Password), nil
}

func (repo *UserRepoImpl) ResetToken(db *gorm.DB, id int) error {
	result := db.
		Model(&domains.Users{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"token":        nil,
			"token_expire": nil,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
