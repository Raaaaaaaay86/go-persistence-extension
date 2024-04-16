package gorme

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ = contract.Basic[any, uint](&BasicRepository[any, uint]{})

type BasicRepository[T any, Q contract.Identifier] struct {
	db *gorm.DB
}

func NewBasicRepository[T any, Q contract.Identifier](db *gorm.DB) *BasicRepository[T, Q] {
	return &BasicRepository[T, Q]{db}
}

func NewEagerBasicRepository[T any, Q contract.Identifier](db *gorm.DB) *BasicRepository[T, Q] {
	return &BasicRepository[T, Q]{db.Preload(clause.Associations)}
}

// Create implements contract.CRUD.
//
// Create a new record.
func (g *BasicRepository[T, Q]) Create(ctx context.Context, entity *T) error {
	return g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(entity).Error
	})
}

// Delete implements contract.CRUD.
// Delete() will look up the primary key of the entity and delete it.
//
// - entity: Describe the entity to be deleted.
func (g *BasicRepository[T, Q]) Delete(ctx context.Context, entity *T) (int64, error) {
	var affectedCount int64

	err := g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if tx := tx.Delete(entity); tx.Error != nil {
			return tx.Error
		} else {
			affectedCount = tx.RowsAffected
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return affectedCount, nil
}

// DeleteById implements contract.CRUD.
func (g *BasicRepository[T, Q]) DeleteById(ctx context.Context, id Q) (int64, error) {
	var entity T
	var affectedCount int64

	err := g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if tx := tx.Delete(&entity, id); tx.Error != nil {
			return tx.Error
		} else {
			affectedCount = tx.RowsAffected
		}
		return nil
	})
	if err != nil {
		return affectedCount, err
	}

	return affectedCount, nil
}

// FindAll implements contract.CRUD.
//
// - limit: -1 means no limit.
//
//	// Return all matched data
//	results, err := FindAll(ctx, -1)
//	// Return matched data with limit with 10
//	results, err := FindAll(ctx, 10)
func (g *BasicRepository[T, Q]) FindAll(ctx context.Context, limit int) ([]*T, error) {
	var results []*T
	if err := g.db.WithContext(ctx).Limit(limit).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// FindBy implements contract.CRUD.
//
// - limit: -1 means no limit.
//
//	// Return all matched data
//	results, err := FindBy(ctx, &user, -1)
//	// Return matched data with limit with 10
//	results, err := FindBy(ctx, &user ,10)
func (g *BasicRepository[T, Q]) FindBy(ctx context.Context, entity T, limit int) ([]*T, error) {
	var results []*T
	if err := g.db.WithContext(ctx).Where(entity).Limit(limit).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// GetBy implements contract.CRUD.
//
// - entity: Describe the entity to be matched.
//
//	// Get first matched user which username is "jordan"
//	result, err := GetBy(ctx, &User{Username: "jordan"})
//	// This will return any first matched user
//	result, err := GetBy(ctx, &User{})
func (g *BasicRepository[T, Q]) GetBy(ctx context.Context, entity T) (*T, error) {
	var result T
	if err := g.db.WithContext(ctx).Where(entity).First(&result).Error; err != nil {
		return &result, err
	}
	return &result, nil
}

// GetById implements contract.CRUD.
//
// - id: Describe the entity's id to be matched.
//
//	// Get user by id equals to 10
//	result, err := GetById(ctx, 10)
func (g *BasicRepository[T, Q]) GetById(ctx context.Context, id Q) (*T, error) {
	var result T
	if err := g.db.First(&result, id).Error; err != nil {
		return &result, err
	}
	return &result, nil
}

// Update implements contract.CRUD.
// Update() will look up the primary key of the entity and update all non-zero fields.
// If the primary key is blank, it will save it as a new record.
func (g *BasicRepository[T, Q]) Update(ctx context.Context, entity *T) (int64, error) {
	var affectedCount int64

	err := g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if tx := tx.Save(entity); tx.Error != nil {
			return tx.Error
		} else {
			affectedCount = tx.RowsAffected
		}
		return nil
	})
	if err != nil {
		return affectedCount, err
	}

	return affectedCount, err
}

// Update implements contract.CRUD.
// Like() will return matched records with like condition.
//
//	// Return all rows which username contains "jordan"
//	results, err := Like(ctx, &User{Username: "jordan"})
func (g *BasicRepository[T, Q]) Like(ctx context.Context, entity T, limit int) ([]*T, error) {
	var results []*T

	fields, err := util.ParseNoneZeroFields(reflect.ValueOf(entity))
	if err != nil {
		return results, err
	}

	db := g.db.WithContext(ctx)

	for _, field := range fields {
		likeStr, ok := field.Value.(string)
		if !ok {
			continue
		}
		db = db.Where(fmt.Sprintf("%s LIKE ?", field.ColumnName), likeStr)
	}

	if err := db.Limit(limit).Find(&results).Error; err != nil {
		return results, err
	}

	return results, nil
}

func (g *BasicRepository[T, Q]) FindTimeBefore(ctx context.Context, entity T, before time.Time, limit int) ([]*T, error) {
	f := fmt.Sprintf
	var results []*T

	field, err := util.ParseTargetField(entity, reflect.TypeOf(time.Time{}))
	if err != nil {
		return results, err
	}

	db := g.db.WithContext(ctx)

	if tx := db.Where(f("%s < ?", field.ColumnName), before).Limit(limit).Find(&results); tx.Error != nil {
		return results, tx.Error
	}

	return results, nil
}

func (g *BasicRepository[T, Q]) FindTimeAfter(ctx context.Context, entity T, after time.Time, limit int) ([]*T, error) {
	f := fmt.Sprintf
	var results []*T

	field, err := util.ParseTargetField(entity, reflect.TypeOf(time.Time{}))
	if err != nil {
		return results, err
	}

	db := g.db.WithContext(ctx)

	if tx := db.Where(f("%s > ?", field.ColumnName), after).Limit(limit).Find(&results); tx.Error != nil {
		return results, tx.Error
	}

	return results, nil
}

func (g *BasicRepository[T, Q]) FindTimeBetween(ctx context.Context, entity T, startAt time.Time, endAt time.Time, limit int) ([]*T, error) {
	f := fmt.Sprintf
	var results []*T

	field, err := util.ParseTargetField(entity, reflect.TypeOf(time.Time{}))
	if err != nil {
		return results, err
	}

	db := g.db.WithContext(ctx)

	tx := db.Where(f("%s BETWEEN ? AND ?", field.ColumnName), startAt, endAt).
		Limit(limit).
		Find(&results)
	if tx.Error != nil {
		return results, tx.Error
	}

	return results, nil
}
