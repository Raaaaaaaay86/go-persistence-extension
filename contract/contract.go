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
	GetBy(ctx context.Context, entity T) (*T, error)
	GetById(ctx context.Context, id Q) (*T, error)
	FindBy(ctx context.Context, entity T, limit int) ([]*T, error)
	FindAll(ctx context.Context, limit int) ([]*T, error)
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) (int64, error)
	Delete(ctx context.Context, entity *T) (int64, error)
	DeleteById(ctx context.Context, id Q) (int64, error)
	Like(ctx context.Context, entity T, limit int) ([]*T, error)
	FindBefore(ctx context.Context, entity T, before time.Time, limit int) ([]*T, error)
	FindAfter(ctx context.Context, entity T, before time.Time, limit int) ([]*T, error)
	FindBetween(ctx context.Context, entity T, startAt time.Time, endAt time.Time, limit int) ([]*T, error)
}

type Paginated[T any, Q Identifier] interface {
	PFindBy(ctx context.Context, entity T, page int, pageSize int) (*Pagination[T], error)
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
