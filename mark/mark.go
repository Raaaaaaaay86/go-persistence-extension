package mark

import (
	"time"

	"gorm.io/gorm"
)

var TargetTime = time.Date(1, 0, 0, 0, 0, 0, 0, time.UTC).Add(1 * time.Nanosecond)

var TargetGormDeleteAt gorm.DeletedAt = gorm.DeletedAt{
	Time:  time.Date(1, 0, 0, 0, 0, 0, 0, time.UTC).Add(1 * time.Nanosecond),
	Valid: true,
}