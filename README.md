# Functional Programming Core Libraries for Go

This repository contains core libraries for functional programming in Go. Below is a description of the packages herein:

- [`fn`](./fn) - functions and structures for working with functions. For example, you can use `Compose` to join two functions together, and `Curry` to split them apart.
- [`iter`](./iter) - functions and structures for working with iterable data, like a slice or `map`. For example, you can iterate a slice, call a function on each element, each in a separate goroutine, and aggregate the result, all from a single function call
- [`slice`](./slice) - fundamental operations on slices. For example, you can get the head (first element) or tail (everything but the first element) of a slice each with a function call.

## Note: Work in Progress

This repository is a work in progress, and is subject to change. This is primarily because it relies heavily on Go's [Generics](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md) feature, which itself is changing. As generics evolves, this library will do so to take advantage of developments in Generics as appropriate.

To use the code herein, you'll need the Go toolchain at least at version 1.17, and you'll need to build with the `-gcflags=-G=3`. For example, to run the tests herein:

```shell
go test -gcflags=-G=3 ./...
```

>At the current version of Go, you'll get an error that looks like the below, and prevents even tests from passing. That's expected and as the generics feature progresses, it will go away.
>
>`internal compiler error: Cannot export a generic function (yet)`.

## Contributing

If you're interested, please contribute any way you wish. Here are some ideas:

- Submit an [issue](https://github.com/go-functional/core/issues) with a question, comment, or anything else
- Make a change and submit a [pull request](https://github.com/go-functional/core/pulls). Even if you aren't a "FP person" or don't know what you're doing, I want to hear what you think, and sometimes the best way to do that is through code :)
- Ping me on Twitter [@arschles](https://twitter.com/arschles)
