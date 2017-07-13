package main

import (
	"log"

	"github.com/go-functional/core/typeclass"
)

func main() {
	iEq1 := typeclass.IntEq(1)
	iEq2 := typeclass.StrEq("1")
	log.Printf("checking if %#v is equal to %#v", iEq1, iEq2)
	log.Printf("the result is: %t", iEq1.Eq(iEq2))

	iSliceEq1 := typeclass.IntSliceEq([]int{1, 2, 3})
	iSliceEq2 := typeclass.IntSliceEq([]int{1, 2, 3})
	log.Printf("checking if %#v is equal to %#v", iSliceEq1, iSliceEq2)
	log.Printf("the result is: %t", iSliceEq1.Eq(iSliceEq2))
}
