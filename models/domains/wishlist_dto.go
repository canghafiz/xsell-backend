package domains

import "time"

type WishlistItem struct {
	WishlistId int       `json:"wishlist_id"`
	UserId     int       `json:"user_id"`
	ProductId  int       `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Product    Products  `json:"product"`
	LikeCount  int64     `json:"like_count"`
}
