package product

import (
	"be/models/domains"
	"time"
)

type SingleProductResource struct {
	ProductId   int               `json:"product_id"`
	Title       string            `json:"title"`
	Slug        string            `json:"slug"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Condition   string            `json:"condition"`
	Status      string            `json:"status"`
	ViewCount   int               `json:"view_count"`
	Category    *CategoryResource `json:"category,omitempty"`
	Images      []ImageResource   `json:"images"`
	Specs       []SpecResource    `json:"specs"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
}

type CategoryResource struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Slug         string `json:"slug"`
}

type ImageResource struct {
	ImageId int    `json:"image_id"`
	URL     string `json:"url"`
}

type SpecResource struct {
	SpecId   int    `json:"spec_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Category string `json:"category,omitempty"` // dari CategoryProductSpec
}

func ToSingleProductResource(product *domains.Products) *SingleProductResource {
	if product == nil {
		return nil
	}

	res := &SingleProductResource{
		ProductId:   product.ProductId,
		Title:       product.Title,
		Slug:        product.ProductSlug,
		Description: product.Description,
		Price:       product.Price,
		Condition:   string(product.Condition),
		Status:      string(product.Status),
		ViewCount:   product.ViewCount,
		CreatedAt:   product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
	}

	if len(product.ProductCategories) > 0 {
		cat := product.ProductCategories[0]
		res.Category = &CategoryResource{
			CategoryId:   cat.CategoryId,
			CategoryName: cat.CategoryName,
			Slug:         cat.CategorySlug,
		}
	}

	// Images
	for _, img := range product.ProductImages {
		res.Images = append(res.Images, ImageResource{
			ImageId: img.ImageId,
			URL:     img.ImageUrl,
		})
	}

	// Specs
	for _, spec := range product.ProductSpecs {
		specRes := SpecResource{
			SpecId: spec.ProductSpecId,
			Name:   spec.CategoryProductSpec.SpecName,
			Value:  spec.SpecValue,
		}
		if spec.CategoryProductSpec.CategoryProductSpecId != 0 {
			specRes.Category = spec.CategoryProductSpec.SpecName
		}
		res.Specs = append(res.Specs, specRes)
	}

	return res
}
