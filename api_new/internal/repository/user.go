package repository

import (
	"context"

	"gorm.io/gorm"
)

type User interface {
	WithContext(ctx context.Context) User
	WithGormDB(db *gorm.DB) User
}

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &user{
		db: db,
	}
}

func (r *user) WithContext(ctx context.Context) User {
	return r
}

func (r *user) WithGormDB(db *gorm.DB) User {
	r.db = db
	return r
}
