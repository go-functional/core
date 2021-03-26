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

// TransformEither runs onLeft if e.Left() == true and onRight otherwise
func ForEither(e Either, onLeft, onRight func() error) error {
	if e.Left() {
		return onLeft()
	}
	return onRight()
}
