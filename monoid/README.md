# Monoids

A Monoid is a functional "container" that has the following operations on it:

- `Zero() T` - returns the zero (in FP, called the "identity") value of the type.
    This is similar in principle to Go's zero values, except you, the creator 
    of the monoid, get to define what the zero value is
- `Append(other T) Monoid<T>` - appends `other` to the callee (the Monoid 
    that you are calling `Append` on ). You get to decide what "append" means.
    For `int`s,  for example, `Append` might be addition, but for `YourType`
    it could mean something completely different
