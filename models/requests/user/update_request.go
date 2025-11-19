package user

import "be/models/domains"

type UpdateDataRequest struct {
	FirstName       string           `json:"first_name" validate:"required,max=25"`
	LastName        *string          `json:"last_name" validate:"max=25"`
	Role            domains.UserRole `json:"role"`
	PhotoProfileUrl *string          `json:"photo_profile_url"`
}

func UpdateDataRequestToDomain(request UpdateDataRequest) domains.Users {
	return domains.Users{
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		Role:            request.Role,
		PhotoProfileUrl: request.PhotoProfileUrl,
	}
}
