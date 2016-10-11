package main

// Equal typeclass:
//  (Eq a) => a -> a -> Bool
type Eq interface {
	Eq(Equaler) bool
}

// Ord typeclass:
//  (Ord a) => a -> a -> Bool
type Ord interface {
	Before(Ord) bool
}

// Show typeclass:
//
//  (Show a) => a -> string
//
// this is the same thing as fmt.Stringer!
type Show interface {
	Show() string
}

// Read typeclass:
//
//  (Read a) => string -> Read
//
// Read a string into a Read implementation
type Read interface {
	Read(string) Read
}

type Either interface {
	Left() bool
	Right() bool
}
