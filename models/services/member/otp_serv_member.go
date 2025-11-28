package member

import (
	"be/models/requests/member/otp"
)

type OtpServMember interface {
	SendOtpToPhoneVerify(request otp.SendRequest) error
	CheckOtpPhoneVerify(request otp.CheckRequest) error
}
