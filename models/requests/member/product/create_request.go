package product

import "be/models/domains"

type CreateProductRequest struct {
	Title         string                   `json:"title" validate:"required,min=3,max=100"`
	Description   string                   `json:"description" validate:"required,min=10,max=5000"`
	Price         float64                  `json:"price" validate:"required,gt=0,lte=999999999"`
	Condition     domains.ProductCondition `json:"condition" validate:"required,oneof=New 'Like New' Good 'Good Quite' 'Needs Repair'"`
	Status        domains.ProductStatus    `json:"status" validate:"required,oneof=Available 'Sold out'"`
	CategoryID    int                      `json:"category_id" validate:"required,gt=0"`
	ListingUserId int                      `json:"listing_user_id" validate:"required,gt=0"`
	Images        []ImageRequest           `json:"images,omitempty"`
	Specs         []SpecRequest            `json:"specs,omitempty"`
	Location      LocationRequest          `json:"location"`
}

func CreateProductRequestToDomain(request CreateProductRequest) *domains.Products {
	return &domains.Products{
		Title:         request.Title,
		Description:   request.Description,
		Price:         request.Price,
		Condition:     request.Condition,
		Status:        request.Status,
		CategoryId:    request.CategoryID,
		ListingUserId: request.ListingUserId,
		ProductImages: ImageRequestToDomains(request.Images),
		ProductSpecs:  SpecRequestToDomains(request.Specs),
		Location:      LocationRequestToDomain(request.Location),
	}
}
