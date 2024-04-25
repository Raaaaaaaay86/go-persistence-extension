package gorme

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/macro"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/macro/operator"
	"github.com/raaaaaaaay86/go-persistence-extension/gorme/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ = contract.Paginated[any, uint](&PaginationRepository[any, uint]{})

type PaginationRepository[T any, Q contract.Identifier] struct {
	db *gorm.DB
}

func NewPaginationRepository[T any, Q contract.Identifier](db *gorm.DB) *PaginationRepository[T, Q] {
	return &PaginationRepository[T, Q]{db}
}

func NewEagerPaginationRepository[T any, Q contract.Identifier](db *gorm.DB) *PaginationRepository[T, Q] {
	return &PaginationRepository[T, Q]{db.Preload(clause.Associations)}
}

// FindAll implements contract.Pagination.
//
// PFindAll() will return all records with pagination.
func (p *PaginationRepository[T, Q]) PFindAll(
	ctx context.Context,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	var results []T
	offset := macro.Offset(page, pageSize)
	if err := p.db.Offset(offset).Limit(pageSize).Find(&results).Error; err != nil {
		return nil, err
	}

	var entity T
	var total int64
	if err := p.db.Model(entity).Count(&total).Error; err != nil {
		return nil, err
	}

	return contract.NewPagination(results, page, pageSize, total), nil
}

// FindBy implements contract.Pagination.
//
// PFindBy() will return matched records with pagination.
func (p *PaginationRepository[T, Q]) PFindBy(
	ctx context.Context,
	query contract.QueryMap,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	var results []T
	if err := p.db.
		Offset(macro.Offset(page, pageSize)).
		Limit(pageSize).
		Where(map[string]interface{}(query)).
		Find(&results).Error; err != nil {
		return nil, err
	}

	var entity T
	var total int64
	if err := p.db.Model(entity).Count(&total).Error; err != nil {
		return nil, err
	}

	return contract.NewPagination(results, page, pageSize, total), nil
}

func (p *PaginationRepository[T, Q]) PFindTimeBefore(
	ctx context.Context,
	entity T,
	before time.Time,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.PFindByTime[T, Q](ctx, p.db, operator.LT, entity, before, page, pageSize)
}

func (p *PaginationRepository[T,
	Q]) PFindTimeAfter(ctx context.Context,
	entity T,
	before time.Time,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.PFindByTime[T, Q](ctx, p.db, operator.GT, entity, before, page, pageSize)
}

func (p *PaginationRepository[T,
	Q]) PFindTimeBetween(ctx context.Context,
	entity T,
	startAt time.Time,
	endAt time.Time,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	f := fmt.Sprintf
	var results []T

	field, err := util.ParseTargetField(entity, reflect.TypeOf(time.Time{}))
	if err != nil {
		return nil, err
	}

	if err := p.db.
		WithContext(ctx).
		Offset(macro.Offset(page, pageSize)).
		Limit(pageSize).
		Where(f("%s > ? AND %s < ?", field.ColumnName, field.ColumnName), startAt, endAt).
		Find(&results).Error; err != nil {
		return nil, err
	}

	total, err := macro.TotalCount[T](ctx, p.db, page, pageSize)
	if err != nil {
		return nil, err
	}

	return contract.NewPagination(results, page, pageSize, total), nil
}

func (p *PaginationRepository[T, Q]) PFindIntGT(
	ctx context.Context,
	entity T,
	value int,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.GT, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindIntGTE(
	ctx context.Context,
	entity T,
	value int,
	page int,
	pageSize int,
) (*contract.Pagination[T],
	error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.GTE, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindIntLT(
	ctx context.Context,
	entity T,
	value int,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.LT, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindIntLTE(
	ctx context.Context,
	entity T,
	value int,
	page int,
	pageSize int,
) (*contract.Pagination[T],
	error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.LTE, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindUintGT(
	ctx context.Context,
	entity T,
	value uint,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.GT, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindUintGTE(
	ctx context.Context,
	entity T,
	value uint,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.GTE, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindUintLT(
	ctx context.Context,
	entity T,
	value uint,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.LT, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindUintLTE(
	ctx context.Context,
	entity T,
	value uint,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.LTE, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindFloat32GT(
	ctx context.Context,
	entity T,
	value float32,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.GT, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindFloat32GTE(
	ctx context.Context,
	entity T,
	value float32,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.GTE, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindFloat32LT(
	ctx context.Context,
	entity T,
	value float32,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.LT, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindFloat32LTE(
	ctx context.Context,
	entity T,
	value float32,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.LTE, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindFloat64GT(
	ctx context.Context,
	entity T,
	value float64,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.GT, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindFloat64GTE(
	ctx context.Context,
	entity T,
	value float64,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.GTE, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindFloat64LT(
	ctx context.Context,
	entity T,
	value float64,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.LT, page, pageSize)
}

func (p *PaginationRepository[T, Q]) PFindFloat64LTE(
	ctx context.Context,
	entity T,
	value float64,
	page int,
	pageSize int,
) (*contract.Pagination[T], error) {
	return macro.ComparePFind(ctx, p.db, entity, value, operator.LTE, page, pageSize)
}
