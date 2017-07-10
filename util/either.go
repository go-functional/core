package util

import (
	"fmt"
)

// Either is a union type - it can be exactly a right or a left, not neither and not both
type Either interface {
	Left() bool
	Right() bool
	fmt.Stringer
}
