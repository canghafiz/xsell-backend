package domains

import "time"

type Otp struct {
	OtpId       int       `gorm:"primary_key;column:otp_id;auto_increment"`
	PhoneNumber string    `gorm:"column:phonenumber"`
	Code        string    `gorm:"column:code"`
	Purpose     string    `gorm:"column:purpose"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	ExpireAt    time.Time `gorm:"column:expire_at;->"`
}

func (Otp) TableName() string {
	return "otp"
}
