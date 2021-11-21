package repository

import (
	"context"

	"github.com/nhannt315/real_estate_api/pkg/errors"

	"github.com/nhannt315/real_estate_api/internal/model"
	"gorm.io/gorm"
)

type User interface {
	WithContext(ctx context.Context) User
	WithGormDB(db *gorm.DB) User
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
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

func (r *user) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, "Find user by email")
	}
	return &user, err
}

func (r *user) Create(user *model.User) error {
	err := r.db.Omit("created_at", "updated_at").Create(user).Error
	if err != nil {
		return errors.Wrap(err, "Create user")
	}
	return nil
}
