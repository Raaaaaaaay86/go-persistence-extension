package gorme

import "github.com/raaaaaaaay86/go-persistence-extension/contract"

type Entity[T contract.Identifier] interface {
	GetID() T
	GetTableName() string
}