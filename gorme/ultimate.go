package gorme

import (
	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ = contract.Ultimate[any, uint](&UltimateRepository[any, uint]{})

type UltimateRepository[T any, Q contract.Identifier] struct {
	contract.Basic[T, Q]
	contract.Paginated[T, Q]
}

func NewUltimateRepository[T any, Q contract.Identifier](
	db *gorm.DB,
) *UltimateRepository[T,Q] {
	return &UltimateRepository[T, Q]{
		NewBasicRepository[T, Q](db),
		NewPaginationRepository[T, Q](db),
	}
}

func NewEagerUltimateRepository[T any, Q contract.Identifier](
	db *gorm.DB,
) *UltimateRepository[T, Q] {
	return NewUltimateRepository[T, Q](db.Preload(clause.Associations))
}
