package util

import (
	"fmt"
)

// Option is the type that can either be something or nothing
type Option interface {
	fmt.Stringer
	Empty() bool
}

// None returns an empty Option
func None() Option {
	return emptyOpt{}
}

type emptyOpt struct{}

func (e emptyOpt) Empty() bool { return true }

func (e emptyOpt) String() string {
	return "empty"
}
