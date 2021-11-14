package model

import "github.com/nhannt315/real_estate_api/pkg/datetime"

type User struct {
	ID uint64 `gorm:"column:id"`

	Email string `gorm:"column:email"`

	PasswordDigest string `gorm:"column:password_digest"`

	CreatedAt *datetime.Time `gorm:"column:created_at"`
	UpdatedAt *datetime.Time `gorm:"column:updated_at"`
}

func (i User) TableName() string {
	return "users"
}
