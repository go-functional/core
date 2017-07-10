package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-functional/core/functor"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	eitherRes := doStuff()
	log.Printf("got back Either<int, error>: %s", eitherRes)

	leftProj := eitherRes.ToLeft()
	rightProj := eitherRes.ToRight()

	log.Printf("left projection (Optional<int>): %s", leftProj)
	log.Printf("right projection (Optional<error>): %s", rightProj)

	leftMapperFunc := func(i int) int {
		return i * 100
	}
	rightMapperFunc := func(e error) error {
		return fmt.Errorf("MAPPED%s", e)
	}

	mappedLeftProj := leftProj.Map(leftMapperFunc)
	mappedRightProj := rightProj.Map(rightMapperFunc)

	log.Printf("mapped left projection (Optional<Int>): %s", mappedLeftProj)
	log.Printf("mapped right projection (<Optional<error>): %s", mappedRightProj)
}

func doStuff() functor.EitherIntOrErr {
	if rand.Int()%2 == 0 {
		return functor.EitherIntOrErrLeft(rand.Int())
	}
	return functor.EitherIntOrErrRight(fmt.Errorf("error%d", rand.Int()))
}
