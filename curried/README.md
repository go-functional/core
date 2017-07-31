# Curried Functions

Known commonly as function closures in Go, curried functions are _very_ important to FP. And they're insanely powerful for Gophers!

## Deconstructing Functions in Go

Did you know that a function with 2 or more arguments can really be broken down into a bunch of functions with 1 argument each??

Take this trivial function:

```go
func Add3(i1, i2, i3 int) int {
  return i1 + i2 + i3
}
```

Below is another _crazy_ way to write the same function:

```go
func Add3(i1 int) func(int) func(int) int { ... }
```

Here it is, put simply, in english: Add3 _closes over_ `i1` and returns a curried function that closes over its argument and returns a similar curried function that finally returns the sum of all 3 parameters.

The implementation, you ask? Here it is!

```go
func Add3(i1 int) func(int) func(int) int {
  return func(i2 int) func(int) int {
    return func(i3 int) int {
      return i1 + i2 + i3
    }
  }
}
```

Hopefully, this looks like recursion - it's very similar!

## Ok, But What About The Real World?

Contrary to what you might believe, curried functions are _not_ as insane as they look in the previous section. We can use currying in the wild as a form of dependency injection. Check this out:

```go
// cl is the dependency of the HandlerFunc (the curried function). We can call getHandler using any cl we like. For example, in unit tests, we can pass an in-memory database!
func getHandler(cl *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // do stuff with cl!
    w.WriteHeader(http.StatusOK)
  }
}
```
