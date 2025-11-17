package services

import "be/models/requests/user"

type UserServ interface {
	RegisterMember(request user.RegisterRequest) error
	LoginMember(request user.LoginRequest) (*string, error)
	Logout(userId int) error
	UpdateData(request user.UpdateDataRequest) (*string, error)
}
