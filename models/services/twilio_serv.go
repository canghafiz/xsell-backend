package services

type TwilioServ interface {
	SendOtpToPhoneNumber(phoneNumber, otpCode string) error
}
