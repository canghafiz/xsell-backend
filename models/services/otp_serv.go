package services

import "be/models/requests/member/otp"

type OtpServ interface {
	SendEmailVerification(request otp.SendRequest) error
	CheckOtp(request otp.CheckRequest) error
}
