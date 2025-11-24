package product

import "be/models/domains"

type UpdateProductRequest struct {
	ProductId   int                       `json:"product_id" validate:"required"`
	Title       *string                   `json:"title,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Price       *float64                  `json:"price,omitempty"`
	Condition   *domains.ProductCondition `json:"condition,omitempty"`
	Status      *domains.ProductStatus    `json:"status,omitempty"`
	CategoryID  *int                      `json:"category_id,omitempty"`
	Images      *[]ImageRequest           `json:"images,omitempty"`
	Specs       *[]SpecRequest            `json:"specs,omitempty"`
}

func UpdateProductRequestToDomain(request UpdateProductRequest) *domains.Products {
	return &domains.Products{
		ProductId:     request.ProductId,
		Title:         *request.Title,
		Description:   *request.Description,
		Price:         *request.Price,
		Condition:     *request.Condition,
		Status:        *request.Status,
		CategoryId:    *request.CategoryID,
		ProductImages: ImageRequestToDomains(*request.Images),
		ProductSpecs:  SpecRequestToDomains(*request.Specs),
	}
}
