package functor

import (
	"log"
)

func intSliceFunctorExample() {
	doSomething := func(i int) int {
		return i + 1
	}
	slice := []int{1, 2, 3, 4}
	results := LiftIntSlice(slice).Map(doSomething).Ints()
	log.Println(results)

	const numInts = 50000
	slice = make([]int, numInts)
	for i := 0; i < numInts; i++ {
		slice[i] = i + 5
	}
	results = LiftIntSlice(slice).Map(doSomething).Ints()
	log.Println(results)
}
