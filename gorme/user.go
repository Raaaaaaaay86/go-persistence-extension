package gorme

import (
	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"gorm.io/gorm"
)


type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type UserRepository struct {
	contract.Basic[*User, uint]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	a := NewBasicRepository[*User, uint](db)
	return &UserRepository{
		Basic: a,
	}
}
