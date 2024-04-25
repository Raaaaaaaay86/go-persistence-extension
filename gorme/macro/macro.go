package macro

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/macro/operator"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/util"
	"gorm.io/gorm"
)

func CompareFind[T any, Q contract.Number](
	ctx context.Context,
	db *gorm.DB,
	entity T,
	value Q,
	operator operator.Enum,
	limit int,
) ([]*T, error) {
	f := fmt.Sprintf
	var results []*T

	var number Q
	field, err := util.ParseTargetField(entity, reflect.TypeOf(number))
	if err != nil {
		return results, err
	}

	db = db.WithContext(ctx)

	tx := db.Where(f("%s %s ?", field.ColumnName, operator), value).
		Limit(limit).
		Find(&results)
	if tx.Error != nil {
		return results, tx.Error
	}

	return results, nil
}

func PFindByTime[T any, Q contract.Identifier](
	ctx context.Context,
	db *gorm.DB,
	operator operator.Enum,
	entity T,
	before time.Time,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	f := fmt.Sprintf
	var results []T

	field, err := util.ParseTargetField(entity, reflect.TypeOf(time.Time{}))
	if err != nil {
		return nil, err
	}

	if err := db.
		Offset(Offset(page, pageSize)).
		Limit(pageSize).
		Where(f("%s %s ?", field.ColumnName, operator), before).
		Find(&results).Error; err != nil {
		return nil, err
	}

	total, err := TotalCount[T](ctx, db, page, pageSize)
	if err != nil {
		return nil, err
	}

	return contract.NewPagination(results, page, pageSize, total), nil
}

func Offset(page int, pageSize int) int {
	return (page - 1) * pageSize
}

func TotalCount[T any](ctx context.Context, db *gorm.DB, page int, pageSize int) (int64, error) {
	var entity T
	var total int64
	if err := db.WithContext(ctx).Model(entity).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
