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

var TargetInt = int(1)
var TargetInt32 = int32(1)
var TargetInt64 = int64(1)

var TargetUint = uint(1)
var TargetUint32 = uint32(1)
var TargetUint64 = uint64(1)

var TargetFloat32 = float32(1)
var TargetFloat64 = float64(1)