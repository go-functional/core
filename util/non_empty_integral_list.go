package util

import (
	"github.com/go-functional/core/typeclass"
)

// NonEmptyIntegralList is a list that has at least one IntegralListElt in it
type NonEmptyIntegralList struct {
	l []typeclass.Integral
}

// NewNonEmptyIntegralList creates a new NonEmptyIntegralList with i and other in it
func NewNonEmptyIntegralList(i typeclass.Integral, other ...typeclass.Integral) NonEmptyIntegralList {
	lst := []typeclass.Integral{i}
	lst = append(lst, other...)
	return NonEmptyIntegralList{l: lst}
}

// Len returns the length of n
func (n NonEmptyIntegralList) Len() int {
	return len(n.l)
}

// Map executes fn against each element of n and returns the resulting list.
//
// This is the functor typeclass implementation
func (n NonEmptyIntegralList) Map(fn func(typeclass.Integral) typeclass.Integral) []typeclass.Integral {
	ret := make([]typeclass.Integral, n.Len())
	for i, elt := range n.l {
		ret[i] = fn(elt)
	}
	return ret
}

// Filter runs fn on every item in n and returns a slice of elements in n for which fn returned true. Elements are returned in order of appearance
func (n NonEmptyIntegralList) Filter(fn func(typeclass.Integral) bool) NonEmptyIntegralList {
	ret := []typeclass.Integral{}
	for _, elt := range n.l {
		if fn(elt) {
			ret = append(ret, elt)
		}
	}
	return NonEmptyIntegralList{l: ret}
}
