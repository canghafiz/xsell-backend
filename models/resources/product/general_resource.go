package product

import "be/models/domains"

type GeneralResource struct {
	ProductId int             `json:"product_id"`
	Title     string          `json:"title"`
	Price     float64         `json:"price"`
	Condition string          `json:"condition"`
	Images    []ImageResource `json:"images"`
}

func ToGeneralResource(product domains.Products) *GeneralResource {
	res := &GeneralResource{
		ProductId: product.ProductId,
		Title:     product.Title,
		Price:     product.Price,
		Condition: string(product.Condition),
	}

	// Images
	for _, img := range product.ProductImages {
		res.Images = append(res.Images, ImageResource{
			ImageId:   img.ImageId,
			URL:       img.ImageUrl,
			OrderSeq:  img.OrderSequence,
			IsPrimary: img.IsPrimary,
		})
	}

	return res
}

func ToGeneralResources(products []domains.Products) []GeneralResource {
	var res []GeneralResource
	for _, product := range products {
		res = append(res, *ToGeneralResource(product))
	}
	return res
}
