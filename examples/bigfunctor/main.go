package main

import (
	"log"

	"github.com/go-functional/core/functor"
)

func main() {
	const numInts = 1000000
	ints := intSlice(numInts)
	log.Printf("lifing %d ints into a functor", len(ints))
	// here's where we turn an int slice into a functor. in FP terms, this is called "lifting"
	functor := functor.LiftIntSlice(ints)
	log.Printf("original functor has %d elements in it", len(functor.Ints()))

	// now we map over the functor to get a new functor with a new int slice in it.

	// this is the mapper function that we will pass to the Map method
	mapperFunc := func(i int) int {
		// note that this func needs to be pure!
		return i + 10
	}

	// The Map func isn't a Go map! Map is a method on all functors that executes mapperFunc
	// on every item in the slice it contains
	//
	// Tip: go into the Map method and make this function execute in serial. See how much slower
	// it is!
	mapped := functor.Map(mapperFunc)
	log.Printf("mapped functor has %d elements in it", len(mapped.Ints()))
}

func intSlice(numInts int) []int {
	slice := make([]int, numInts)
	for i := 0; i < numInts; i++ {
		slice[i] = i + 5
	}
	return slice
}
