package product

import "be/models/domains"

type LocationRequest struct {
	UserId    int     `json:"user_id" validate:"required,gt=0"`
	Latitude  float64 `json:"latitude" validate:"required,gt=-90,lte=90"`
	Longitude float64 `json:"longitude" validate:"required,gt=-180,lte=180"`
}

func LocationRequestToDomain(request LocationRequest) domains.Location {
	return domains.Location{
		UserId:    request.UserId,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
	}
}
