package entity

import (
	"time"

	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Age      int
	Birthday time.Time
}

type UserQueryMapper struct {
	ID        *uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Username  *string
	Email     *string
	Age       *int
	Birthday  *time.Time
}

func (u UserQueryMapper) ToMap() contract.QueryMap {
	return gorme.ToQueryMap(u)
}