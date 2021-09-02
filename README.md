# Functional Programming Core Libraries for Go

This repository contains core libraries for functional programming (FP) in Go. Below is a description of the packages herein:

- [`fn`](./fn) - functions and structures for working with functions. For example, you can use `Compose` to join two functions together, and `Curry` to split them apart.
- [`slice`](./slice) - fundamental operations on slices. For example, you can get the head (first element) or tail (everything but the first element) of a slice each with a function call.

## FP Theory

This repository aims to draw on FP theory to introduce useful functionality to the Go ecosystem. As such, you'll likely not see data structures or containers like `Option`, `Functor`, etc... here. You will, however, see the practical applications of the functionality those concepts. \

For example, this repository provides convenient `Map` and `FlatMap` functionality over slices, which would generally be under the purview of functors and monads, respectively.

## Note: Work in Progress

This repository is a work in progress, and is subject to change. This is primarily because it relies heavily on Go's [Generics](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md) feature, which itself is changing. As generics evolves, this library will do so to take advantage of developments in Generics as appropriate.

To use the code herein, you'll need the Go toolchain at least at version 1.17, and you'll need to build with the `-gcflags=-G=3`. For example, to run the tests herein:

```shell
go test -gcflags=-G=3 ./...
```

>At the current version of Go, you'll get an error that looks like the below, and prevents even tests from passing. That's expected and as the generics feature progresses, it will go away.
>
>`internal compiler error: Cannot export a generic function (yet)`.

## Resources and Learning

Not many resources exist for learning how to apply FP concepts to Go. To learn more about the concepts herein, please see the following:

- [FP explained with Javascript](https://github.com/hemanth/functional-programming-jargon)
- [Learn You a Haskell for Great Good](http://learnyouahaskell.com/)

## Contributing

If you're interested, please contribute any way you wish. Here are some ideas:

- Submit an [issue](https://github.com/go-functional/core/issues) with a question, comment, or anything else
- Make a change and submit a [pull request](https://github.com/go-functional/core/pulls). Even if you aren't a "FP person" or don't know what you're doing, I want to hear what you think, and sometimes the best way to do that is through code :)
- Ping me on Twitter [@arschles](https://twitter.com/arschles)
