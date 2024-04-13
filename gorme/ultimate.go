package gorme

import (
	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ = contract.Ultimate[Entity[uint], uint](&UltimateRepository[Entity[uint], uint]{})

type UltimateRepository[T any, Q contract.Identifier] struct {
	contract.Basic[T, Q]
	contract.Paginated[T, Q]
}

func NewUltimateRepository(
	db *gorm.DB,
) *UltimateRepository[Entity[uint], uint] {
	return &UltimateRepository[Entity[uint], uint]{
		NewBasicRepository(db),
		NewPaginationRepository(db),
	}
}

func NewEagerUltimateRepository(
	db *gorm.DB,
) *UltimateRepository[Entity[uint], uint] {
	return NewUltimateRepository(db.Preload(clause.Associations))
}
