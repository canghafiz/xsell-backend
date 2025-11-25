package member

import (
	"be/models/requests/member/product"
	res "be/models/resources/product"
)

type ProductServMember interface {
	Create(request product.CreateProductRequest) error
	Update(request product.UpdateProductRequest) error
	GetSingleBySlug(productSlug string) (*res.SingleProductResource, error)
	Delete(productId int) error
}
