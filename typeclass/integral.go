package typeclass

import (
	"strconv"
)

// Integral is the typeclass for representing an integer type
type Integral interface {
	Int() int64
}

type IntIntegral int

func (i IntIntegral) Int() int64 {
	return int64(i)
}

type Int32Integral int32

func (i Int32Integral) Int() int64 {
	return int64(i)
}

type Int64Integral int64

func (i Int64Integral) Int() int64 {
	return int64(i)
}

type StringIntegral string

func (s StringIntegral) Int() int64 {
	i, err := strconv.Atoi(string(s))
	if err != nil {
		return 0
	}
	return int64(i)
}
