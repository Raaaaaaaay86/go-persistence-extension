package gorme

import (
	"context"

	"github.com/raaaaaaaay86/go-persistence-extension/contract"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var _ = contract.Paginated[GormEntity[uint], uint](&PaginationRepository[GormEntity[uint], uint]{})

type PaginationRepository[T any, Q contract.Identifier] struct {
	db *gorm.DB
}

func NewPaginationRepository(db *gorm.DB) *PaginationRepository[GormEntity[uint], uint] {
	return &PaginationRepository[GormEntity[uint], uint]{db}
}

func NewEagerPaginationRepository(db *gorm.DB) *PaginationRepository[GormEntity[uint], uint] {
	return &PaginationRepository[GormEntity[uint], uint]{db.Preload(clause.Associations)}
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
	offset := p.offset(page, pageSize)
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
	entity T, 
	page int, 
	pageSize int,
) (*contract.Pagination[T], error) {
	var results []T
	offset := p.offset(page, pageSize)
	if err := p.db.Offset(offset).Limit(pageSize).Where(entity).Find(&results).Error; err != nil {
		return nil, err
	}

	var total int64
	if err := p.db.Model(entity).Count(&total).Error; err != nil {
		return nil, err
	}

	return contract.NewPagination(results, page, pageSize, total), nil
}

func (p *PaginationRepository[T, Q]) offset(page int, pageSize int) int {
	return (page - 1) * pageSize
}
