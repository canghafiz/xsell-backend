package otp

import "be/models/domains"

type CheckRequest struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}

func CheckRequestToDomains(request CheckRequest) domains.Otp {
	return domains.Otp{
		PhoneNumber: request.PhoneNumber,
		Code:        request.Code,
	}
}
