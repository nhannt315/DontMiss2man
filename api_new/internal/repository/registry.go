package repository

import "gorm.io/gorm"

type Registry interface {
	UserRepository() User
}

type registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{
		db: db,
	}
}

func (r *registry) UserRepository() User {
	return NewUserRepository(r.db)
}
