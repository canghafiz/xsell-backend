package user

import (
	"be/helper"
	"be/models/domains"
)

type RegisterRequest struct {
	FirstName string  `json:"first_name" validate:"required,max=25"`
	LastName  *string `json:"last_name" validate:"max=25"`
	Email     string  `json:"email" validate:"required,email,max=100"`
	Password  string  `json:"password" validate:"required,min=6,max=100"`
}

func RegisterRequestToDomain(request RegisterRequest) domains.Users {
	return domains.Users{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  helper.HashedPassword(request.Password),
	}
}
