package contract

import "golang.org/x/exp/constraints"

type Identifier interface {
	string| constraints.Integer | constraints.Signed| constraints.Unsigned
}