package main

import (
	"log"

	"github.com/go-functional/core/functor"
)

func main() {
	intSlice := []int{1, 2, 3, 4}
	log.Printf("lifing this int slice into a functor: %#v", intSlice)
	// here's where we turn an int slice into a functor. in FP terms, this is called "lifting"
	functor := functor.LiftIntSlice(intSlice)
	log.Printf("original functor: %s", functor)

	// now we map over the functor to get a new functor with a new int slice in it.

	// this is the mapper function that we will pass to the Map method
	mapperFunc := func(i int) int {
		// note that this func needs to be pure!
		return i + 10
	}

	// The Map func isn't a Go map! Map is a method on all functors that executes mapperFunc
	// on every item in the slice it contains
	mapped := functor.Map(mapperFunc)
	log.Printf("mapped functor: %s", mapped)
}
