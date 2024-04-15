package repository

import (
	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	contract.Ultimate[entity.User, uint]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
		Ultimate: gorme.NewEagerUltimateRepository[entity.User, uint](db),
	}
}
