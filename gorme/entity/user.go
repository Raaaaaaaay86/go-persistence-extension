package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Age      int
	Birthday time.Time
}
