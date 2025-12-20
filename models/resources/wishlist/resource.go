package wishlist

import (
	"be/models/domains"
	"time"
)

type Resource struct {
	WishlistId int             `json:"wishlist_id"`
	Product    ProductResource `json:"product"`
}

func ToResource(wishlist domains.WishlistItem) Resource {
	return Resource{
		WishlistId: wishlist.WishlistId,
		Product:    ToProductResource(wishlist.Product, int(wishlist.LikeCount)),
	}
}

func ToResources(wishlists []domains.WishlistItem) []Resource {
	var items []Resource
	for _, wishlist := range wishlists {
		items = append(items, ToResource(wishlist))
	}
	return items
}

type ProductResource struct {
	ProductId int                   `json:"product_id"`
	Title     string                `json:"title"`
	ViewCount int                   `json:"view_count"`
	MainImage string                `json:"main_image"`
	TotalLike int                   `json:"total_like"`
	Price     float64               `json:"price"`
	Status    domains.ProductStatus `json:"status"`
	CreatedAt time.Time             `json:"created_at"`
}

func ToProductResource(product domains.Products, totalLike int) ProductResource {
	mainImage := product.ProductImages[0].ImageUrl

	return ProductResource{
		ProductId: product.ProductId,
		Title:     product.Title,
		ViewCount: product.ViewCount,
		Price:     product.Price,
		Status:    product.Status,
		MainImage: mainImage,
		TotalLike: totalLike,
		CreatedAt: product.CreatedAt,
	}
}
