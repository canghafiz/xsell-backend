package impl

import (
	"be/models/domains"
	"errors"

	"gorm.io/gorm"
)

type WishlistRepoMemberImpl struct {
}

func NewWishlistRepoMemberImpl() *WishlistRepoMemberImpl {
	return &WishlistRepoMemberImpl{}
}

func (repo *WishlistRepoMemberImpl) UpdateWishlist(db *gorm.DB, wishlist domains.Wishlist) error {
	var existing domains.Wishlist
	err := db.Where("user_id = ? AND product_id = ?", wishlist.UserId, wishlist.ProductId).First(&existing).Error

	if errors.Is(gorm.ErrRecordNotFound, err) {
		// Not found → create new
		return db.Create(&wishlist).Error
	}

	if err != nil {
		// Other DB error
		return err
	}

	// Found → delete
	return db.Delete(&existing).Error
}

func (repo *WishlistRepoMemberImpl) GetWishlistsByUserId(db *gorm.DB, userID int) ([]domains.WishlistItem, error) {
	var items []domains.WishlistItem

	err := db.Raw(`
		SELECT 
			w.wishlist_id,
			w.user_id,
			w.product_id,
			w.created_at,
			w.updated_at,
			p.product_id        AS "product.product_id",
			p.listing_user_id   AS "product.listing_user_id",
			p.sub_category_id   AS "product.sub_category_id",
			p.title             AS "product.title",
			p.product_slug      AS "product.product_slug",
			p.description       AS "product.description",
			p.price             AS "product.price",
			p.condition         AS "product.condition",
			p.status            AS "product.status",
			p.view_count        AS "product.view_count",
			p.created_at        AS "product.created_at",
			p.updated_at        AS "product.updated_at",
			(SELECT COUNT(*) FROM wishlists w2 WHERE w2.product_id = w.product_id) AS like_count
		FROM wishlists w
		INNER JOIN products p ON w.product_id = p.product_id
		WHERE w.user_id = ?
		ORDER BY w.created_at DESC
	`, userID).Scan(&items).Error

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *WishlistRepoMemberImpl) CheckWishlist(db *gorm.DB, wishlist domains.Wishlist) (bool, error) {
	var count int64
	err := db.
		Model(&domains.Wishlist{}).
		Where("user_id = ? AND product_id = ?", wishlist.UserId, wishlist.ProductId).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
