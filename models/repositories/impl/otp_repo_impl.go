package impl

import (
	"be/models/domains"
	"fmt"

	"gorm.io/gorm"
)

type OtpRepoImpl struct {
}

func NewOtpRepoImpl() *OtpRepoImpl {
	return &OtpRepoImpl{}
}

func (repo *OtpRepoImpl) SendOtp(db *gorm.DB, otp domains.Otp) error {
	err := db.Create(&otp).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *OtpRepoImpl) CheckOtp(db *gorm.DB, otp domains.Otp) (bool, error) {
	var storedOtp domains.Otp

	err := db.
		Where("phone_number = ? AND code = ? AND purpose = ?", otp.PhoneNumber, otp.Code, otp.Purpose).
		Where("expire_at > NOW()").
		First(&storedOtp).Error

	if err != nil {
		return false, err
	}

	if err := db.Delete(&storedOtp).Error; err != nil {
		return false, fmt.Errorf("failed to delete OTP: %w", err)
	}

	return true, nil
}
