package gorme

import "github.com/raaaaaaaay86/go-persistence-extension/contract"

type GormEntity[T contract.Identifier] interface {
	GetID() T
	GetTableName() T
}