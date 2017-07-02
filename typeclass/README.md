# Typeclasses

The type class is a foundational and powerful FP concept. Similar to `interface`s in Go, type classes allow you to define some behavior on a type, but not implement it.

Since the implementation is done elsewhere, this is a form of polymorphism, but much more flexible or "ad-hoc" than the traditional kind in OOP languages, and a bit more than even standard composition in Go.

Type classes, are different from standard `interface`s because the implementation of the type is separate from the concrete type.

## Warning

Type classes in Go don't fit into the language very well, and here's why:

1. Go doesn't have generics or classical higher-order functions
2. The Go compiler doesn't have the ability to implicitly convert one type to another

## So, Why Are We Here?

Go makes it hard, but type classes are a powerful tool for us in specicic situations. We'll use examples hereafter to explain the utility of this pattern.

## Checking Equality

Let's say we have a type that looks like this:

```go
type MyType struct {
  I int
}
```

How should we check if one `MyType` value is equal to another? We could start with this:

```go
func (m MyType) Eq(other MyType) bool {
  return m.I == other.I // or some more complicated logic!
}
```

## Comparing Across Types

The above `Equals` method will get us pretty far, but what if you want to compare `int`s with your `MyType` value? Let's use a type switch to do that:

```go
func (m MyType) Eq(other interface{}) bool {
  switch x := other.(type) {
  case MyType:
    return x.I == m.I
  case int:
    return m.I == x
  default:
    return false
  }
}
```

## When You Don't Own the Code

Let's take it a step further now. Notice how `Equals` is defined directly on `MyType`. If you don't "own" the code that `MyType` is written in, you're out of luck if it doesn't have `Equals` (or similar) defined on it. We can get over that by writing up an adapter:

```go
type MyTypeEqualer pkg.MyType

func (m MyTypeEqualer) Eq(other interface{}) bool {
  mt := pkg.MyType(m)
  switch x := other.(type) {
  case pkg.MyType:
    return mt.I == x.I
  case int:
    return mt.I == x
  default:
    return false
  }
}
```

Now, we can use the adapter anywhere we want:

```go
i := 2
mt := pkg.MyType{I: 2}
isEq := MyTypeEqualer(mt).Eq(i)
// isEq is true!
```

When we called `MyTypeEqualer(mt)`, we converted our `MyType` value to the wrapper type. We call this _applying_ the type class. Other languages will automatically do this application for us, but in Go we have to do it manually.

## Generalizing

Let's take it one final step and write an `interface` so that we can:

1. Check if _any value of any type_ is equal to any other
2. Generically refer to things that have an `Eq` method on them

Here's roughly what it would look like:

```go
type Eq interface {
  Eq(Eq) bool
}
```

Now, we can write very generic functions like this:

```go
// AllEqual returns true if eq is equal to all of eqs, false otherwise
func AllEqual(eq, eqs ...Eq) bool {
  for _, e := range eqs {
    if !eq.Eq(e) {
      return false
    }
  }
  return true
}
```

And finally, we can check if lots of different types are _semantically_ equal:

```go
mt := MyType{I: 1}
// IntEqualer and StringEqualer are omitted, but have similar implementations as we've seen in this document
allAreEqual := AllEqual(MyTypeEqualer(mt), IntEqualer(1), StringEqualer("1"))
```
