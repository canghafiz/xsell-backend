package member

import "be/models/requests/user"

type UserServMember interface {
	Register(request user.RegisterRequest) error
	Login(request user.LoginRequest) (*string, error)
	Logout(userId int) error
	UpdateData(request user.UpdateDataRequest) (*string, error)
}
