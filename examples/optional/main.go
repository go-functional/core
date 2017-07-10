package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-functional/core/functor"
)

func main() {
	const i = 123

	// create empty and full Option<Int>
	emptyInt := functor.EmptyInt()
	fullInt := functor.SomeInt(i)
	log.Printf("created OptionalInts %s and %s", emptyInt, fullInt)

	// now map over them
	intMapperFunc := func(i int) int {
		return i * 100
	}
	mappedEmptyInt := emptyInt.Map(intMapperFunc)
	mappedFullInt := fullInt.Map(intMapperFunc)

	log.Printf("mapped OptionalInts %s and %s", mappedEmptyInt, mappedFullInt)

	// create empty and full Option<error>
	emptyErr := functor.EmptyErr()
	fullErr := functor.SomeErr(errors.New("test error"))
	log.Printf("created OptionalErrs %s and %s", emptyErr, fullErr)

	// now map over them
	errMapperFunc := func(e error) error {
		return fmt.Errorf("MAPPED %s", e)
	}

	mappedEmptyErr := emptyErr.Map(errMapperFunc)
	mappedFullErr := fullErr.Map(errMapperFunc)
	log.Printf("mapped OptionalErrs %s and %s", mappedEmptyErr, mappedFullErr)
}
