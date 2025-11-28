package domains

type UserVerified struct {
	VerifiedId int    `gorm:"primary_key;column:verified_id;auto_increment"`
	UserId     int    `gorm:"column:user_id"`
	Email      string `gorm:"column:email"`
	Verified   bool   `gorm:"column:verified"`
	CreatedAt  string `gorm:"column:created_at"`
	UpdatedAt  string `gorm:"column:updated_at"`
}

func (UserVerified) TableName() string {
	return "userverified"
}
