package services

import "be/models/requests/user"

type UserServ interface {
	ChangePw(request user.ChangePwRequest, jwtValue string) error
}
