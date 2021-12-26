package models

import (
	"errors"

	"gorm.io/gorm"
)

// UserRepo implements models.UserRepository
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo ..
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

// FindByID ..FindByToken(token string) (*User, error)
func (r *UserRepo) FindByToken(token string) (*User, error) {
	var user = User{}

	if err := r.db.Table("users").First(&user, "token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Can't find token of user in database")
		} else {
			return nil, err
		}
	}
	return &user, nil
}
