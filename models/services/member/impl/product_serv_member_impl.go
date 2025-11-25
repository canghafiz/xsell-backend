package impl

import (
	"be/helpers"
	"be/models/domains"
	"be/models/repositories/member"
	"be/models/requests/member/product"
	res "be/models/resources/product"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServMemberImpl struct {
	Db          *gorm.DB
	Validator   *validator.Validate
	ProductRepo member.ProductRepoMember
}

func NewProductServMemberImpl(db *gorm.DB, validator *validator.Validate, productRepo member.ProductRepoMember) *ProductServMemberImpl {
	return &ProductServMemberImpl{Db: db, Validator: validator, ProductRepo: productRepo}
}

func (serv *ProductServMemberImpl) Create(request product.CreateProductRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	model := product.CreateProductRequestToDomain(request)

	// Call repo
	err := serv.ProductRepo.Create(serv.Db, *model)
	if err != nil {
		log.Printf("[ProductRepoMember.Create] error: %v", err)
		return fmt.Errorf("failed to create product, please try again later")
	}

	return nil
}

func (serv *ProductServMemberImpl) Update(request product.UpdateProductRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	model := product.UpdateProductRequestToDomain(request)

	// Call repo
	err := serv.ProductRepo.Update(serv.Db, *model)
	if err != nil {
		log.Printf("[ProductRepoMember.Update] error: %v", err)
		return fmt.Errorf("failed to update product, please try again later")
	}

	return nil
}

func (serv *ProductServMemberImpl) GetSingleBySlug(productSlug string) (*res.SingleProductResource, error) {
	// Call repo
	result, err := serv.ProductRepo.GetSingleBySlug(serv.Db, productSlug)
	if err != nil {
		log.Printf("[ProductRepoMember.GetSingleBySlug] error: %v", err)
		return nil, fmt.Errorf("failed to get single product by slug, please try again later")
	}

	return res.ToSingleProductResource(result), nil
}

func (serv *ProductServMemberImpl) Delete(productId int) error {
	// Call repo
	err := serv.ProductRepo.Delete(serv.Db, domains.Products{
		ProductId: productId,
	})
	if err != nil {
		log.Printf("[ProductRepoMember.Delete] error: %v", err)
		return fmt.Errorf("failed to delete by product id, please try again later")
	}

	return nil
}
