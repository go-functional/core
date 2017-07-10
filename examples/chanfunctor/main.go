package main

import (
	"log"

	"github.com/go-functional/core/functor"
)

func main() {
	const numInts = 1000000

	// start a goroutine to feed 1,000,000 ints into the int channel. The functor
	// below will consume from the int channel when we call map on it.
	intCh := make(chan int)
	go func() {
		ints := intSlice(numInts)
		log.Printf("feeding %d integers from a goroutine into the int channel that the functor consumes", len(ints))
		for _, elt := range ints {
			intCh <- elt
		}
		// We have to close the channel, or otherwise the Map method will never return
		close(intCh)
	}()

	// lift the int channel into a functor
	log.Printf("lifting int channel into a functor")
	functor := functor.LiftIntSliceFromChan(intCh)
	log.Printf("initial functor has %d elements in it", len(functor.Ints()))
	log.Printf("this is because none of the integers have been fed to it via the int channel!")

	/////
	// now we map over the functor to get a new functor with a new int slice in it.
	/////

	// this is the mapper function that we will pass to the Map method
	mapperFunc := func(i int) int {
		// note that this func needs to be pure!
		return i + 10
	}

	// The Map func isn't a Go map! Map is a method on all functors that executes mapperFunc
	// on every item in the slice it contains
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
