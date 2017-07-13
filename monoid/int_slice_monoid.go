package monoid

// IntSliceMonoid is a monoid container for int slices
type IntSliceMonoid interface {
	// Zero returns the "zero value" of an int slice, which is nil
	Zero() []int
	// Append appends to the contained int slice
	Append(i ...int) IntSliceMonoid
	// Ints is the "escape hatch" for this monoid, to integrate with non-FP
	// Go code
	Ints() []int
}

// LiftIntSlice lifts an int slice into an IntSliceMonoid
func LiftIntSlice(ints []int) IntSliceMonoid {
	return intSliceMonoid{ints: ints}
}

type intSliceMonoid struct {
	ints []int
}

func (intSliceMonoid) Zero() []int {
	return nil
}

func (i intSliceMonoid) Append(ints ...int) IntSliceMonoid {
	i.ints = append(i.ints, ints...)
	return i
}

func (i intSliceMonoid) Ints() []int {
	return i.ints
}
