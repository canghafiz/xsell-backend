package impl

import (
	"be/helpers"
	"be/models/repositories"
	"be/models/requests/member/otp"
	"be/models/services"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OtpServMemberImpl struct {
	Db         *gorm.DB
	Validator  *validator.Validate
	OtpRepo    repositories.OtpRepo
	TwilioServ services.TwilioServ
}

func NewOtpServMemberImpl(db *gorm.DB, validator *validator.Validate, otpRepo repositories.OtpRepo, twilioServ services.TwilioServ) *OtpServMemberImpl {
	return &OtpServMemberImpl{Db: db, Validator: validator, OtpRepo: otpRepo, TwilioServ: twilioServ}
}

func (serv *OtpServMemberImpl) SendOtpToPhoneVerify(request otp.SendRequest) error {
	if err := helpers.ErrValidator(request, serv.Validator); err != nil {
		return err
	}

	tx := serv.Db.Begin()
	if tx.Error != nil {
		log.Printf("Error starting transaction: %v", tx.Error)
		return fmt.Errorf("failed to send OTP, please try again later")
	}

	// Auto-rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Rollback due to panic: %v", r)
		} else if tx.Error != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back: %v", tx.Error)
		}
	}()

	// Generate model
	model := otp.SendRequestToDomain(request)
	code, err := serv.generateNumericOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP")
	}
	model.Code = code

	//  Call Repo
	if err := serv.OtpRepo.SendOtp(tx, model); err != nil {
		log.Printf("[OtpRepo.SendOtp] error: %v", err)
		return fmt.Errorf("failed to send OTP, please try again later")
	}

	// Call Twilio
	if err := serv.TwilioServ.SendOtpToPhoneNumber(model.PhoneNumber, model.Code); err != nil {
		log.Printf("[TwilioServ.SendOtpToPhoneNumber] error: %v", err)
		return fmt.Errorf("OTP generated but failed to deliver, please try again")
	}

	// Commit
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return fmt.Errorf("OTP sent but failed to finalize, please contact support")
	}

	return nil
}

func (serv *OtpServMemberImpl) CheckOtpPhoneVerify(request otp.CheckRequest) error {
	if err := helpers.ErrValidator(request, serv.Validator); err != nil {
		return err
	}

	model := otp.CheckRequestToDomains(request)
	model.Purpose = "phone_verification"

	// Call repo
	result, err := serv.OtpRepo.CheckOtp(serv.Db, model)
	if err != nil {
		log.Printf("[OtpRepo.CheckOtp] error: %v", err)
		return fmt.Errorf("failed to check OTP, please try again later")
	}
	if !result {
		log.Printf("Otp repo result is false")
		return fmt.Errorf("failed to check OTP, please try again later")
	}

	return nil
}

func (serv *OtpServMemberImpl) generateNumericOTP(length int) (string, error) {
	const digits = "0123456789"
	result := make([]byte, length)
	maxCode := big.NewInt(int64(len(digits)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, maxCode)
		if err != nil {
			return "", err
		}
		result[i] = digits[n.Int64()]
	}

	return string(result), nil
}
