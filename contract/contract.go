package contract

import (
	"context"
	"time"
)

type Ultimate[T any, Q Identifier] interface {
	Basic[T, Q]
	Paginated[T, Q]
}

type Basic[T any, Q Identifier] interface {
	GetBy(ctx context.Context, query QueryMap) (*T, error)
	GetById(ctx context.Context, id Q) (*T, error)
	FindBy(ctx context.Context, query QueryMap, limit int) ([]*T, error)
	FindAll(ctx context.Context, limit int) ([]*T, error)
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) (int64, error)
	Delete(ctx context.Context, entity *T) (int64, error)
	DeleteById(ctx context.Context, id Q) (int64, error)
	Like(ctx context.Context, entity T, limit int) ([]*T, error)
	FindTimeBefore(ctx context.Context, entity T, before time.Time, limit int) ([]*T, error)
	FindTimeAfter(ctx context.Context, entity T, before time.Time, limit int) ([]*T, error)
	FindTimeBetween(ctx context.Context, entity T, startAt time.Time, endAt time.Time, limit int) ([]*T, error)
	FindIntGT(ctx context.Context, entity T, value int, limit int) ([]*T, error)
	FindIntGTE(ctx context.Context, entity T, value int, limit int) ([]*T, error)
	FindIntLT(ctx context.Context, entity T, value int, limit int) ([]*T, error)
	FindIntLTE(ctx context.Context, entity T, value int, limit int) ([]*T, error)
	FindUintGT(ctx context.Context, entity T, value uint, limit int) ([]*T, error)
	FindUintGTE(ctx context.Context, entity T, value uint, limit int) ([]*T, error)
	FindUintLT(ctx context.Context, entity T, value uint, limit int) ([]*T, error)
	FindUintLTE(ctx context.Context, entity T, value uint, limit int) ([]*T, error)
	FindFloat32GT(ctx context.Context, entity T, value float32, limit int) ([]*T, error)
	FindFloat32GTE(ctx context.Context, entity T, value float32, limit int) ([]*T, error)
	FindFloat32LT(ctx context.Context, entity T, value float32, limit int) ([]*T, error)
	FindFloat32LTE(ctx context.Context, entity T, value float32, limit int) ([]*T, error)
	FindFloat64GT(ctx context.Context, entity T, value float64, limit int) ([]*T, error)
	FindFloat64GTE(ctx context.Context, entity T, value float64, limit int) ([]*T, error)
	FindFloat64LT(ctx context.Context, entity T, value float64, limit int) ([]*T, error)
	FindFloat64LTE(ctx context.Context, entity T, value float64, limit int) ([]*T, error)
}

type Paginated[T any, Q Identifier] interface {
	PFindBy(ctx context.Context, query QueryMap, page int, pageSize int) (*Pagination[T], error)
	PFindAll(ctx context.Context, page int, pageSize int) (*Pagination[T], error)
}

type Pagination[T any] struct {
	Page       int
	PageSize   int
	TotalPage  int
	TotalCount int64
	Results    []T
}

func NewPagination[T any](results []T,page int, size int, total int64) *Pagination[T] {
	return &Pagination[T]{
		Page:       page,
		PageSize:   size,
		TotalPage:  int(total) / size,
		TotalCount: total,
		Results:    results,
	}
}

type QueryMap map[string]interface{}