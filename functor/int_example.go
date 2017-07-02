package functor

import (
	"log"
)

func intSliceFunctorExample() {
	doSomething := func(i int) int {
		return i + 1
	}
	slice := []int{1, 2, 3, 4}
	results := ToIntSliceFunctor(slice).Map(doSomething).Ints()
	log.Println(results)
}
