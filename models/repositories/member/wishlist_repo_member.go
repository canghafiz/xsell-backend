package member

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type WishlistRepoMember interface {
	UpdateWishlist(db *gorm.DB, wishlist domains.Wishlist) error
	GetWishlistsByUserId(db *gorm.DB, userId int) ([]domains.WishlistItem, error)
	CheckWishlist(db *gorm.DB, wishlist domains.Wishlist) (bool, error)
}
