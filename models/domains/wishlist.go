package domains

import "time"

type Wishlist struct {
	WishlistId int       `gorm:"primary_key;column:wishlist_id;auto_increment"`
	UserId     int       `gorm:"column:user_id"`
	ProductId  int       `gorm:"column:product_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	Product    Products  `gorm:"foreignKey:product_id;references:product_id"`
}

func (Wishlist) TableName() string {
	return "wishlists"
}
