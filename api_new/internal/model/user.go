package model

import "github.com/nhannt315/real_estate_api/pkg/datetime"

type User struct {
	ID uint64 `gorm:"column:id"`

	Email string `gorm:"column:email"`

	PasswordDigest string `gorm:"column:password_digest"`

	AccessToken string `gorm:"-"`

	CreatedAt datetime.NullTime `gorm:"column:created_at"`
	UpdatedAt datetime.NullTime `gorm:"column:updated_at"`
}

func (i User) TableName() string {
	return "users"
}
