package wishlist

type CheckRequest struct {
	UserId    int `json:"user_id" validate:"required"`
	ProductId int `json:"product_id" validate:"required"`
}
