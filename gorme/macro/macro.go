package macro

import (
	"context"
	"fmt"
	"reflect"

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
