package contract

import "golang.org/x/exp/constraints"

type Identifier interface {
	string| constraints.Integer | constraints.Signed| constraints.Unsigned
}

type Number interface {
	constraints.Integer | constraints.Signed| constraints.Unsigned | constraints.Float
}