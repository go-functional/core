package functor

import (
	"log"
)

func intSliceFunctor() {
	doSomething := func(i int) int {
		return i + 1
	}
	slice := []int{1, 2, 3, 4}
	results := IntSliceFunctor{ints: slice}.Map(doSomething).Ints()
	log.Println(results)
}
