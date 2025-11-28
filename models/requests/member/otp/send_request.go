package otp

import "be/models/domains"

type SendRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,numeric,len=11"`
}

func SendRequestToDomain(request SendRequest) domains.Otp {
	return domains.Otp{
		PhoneNumber: request.PhoneNumber,
		Purpose:     "phone_verification",
	}
}
