package par

// Tree is a parallel execution binary tree.
// Each node in this binary tree has three functions
// defined on it, and a single operation that it can
// do called Exec. The three functions are called
// left, right and reduce.
// When reduce is called, left and right run to
// completion each in their own goroutines.
// Then, when both are complete, the values are passed
// to reduce and a final value is returned from that function.
//
// The tree can be have any number of nodes and can
// be as balanced or unbalanced as you decide.
// For any given tree node, add left or right nodes
// to it via the AddLeft and AddRight functions,
// respectively.
//
// Parallel execution binary trees are intended for problems
// in which work can be broken down into multiple dependent
// pieces. For example, if you have a long list of files
// to process, need to return a single file with the result
// of the processing operation, and can process each
// file completely separately from the others, this model
// may work for you.
type Tree[L, R, T any] struct {
	reduce func(L, R) T
	left   func() L
	right  func() R
}

// Exec runs the tree and returns the result.
func (t *Tree[L, R, T]) Exec() T {
	leftChan := make(chan L)
	rightChan := make(chan R)
	go func() {
		leftChan <- t.left()
	}()
	go func() {
		rightChan <- t.right()
	}()

	leftVal := <-leftChan
	rightVal := <-rightChan
	return t.reduce(leftVal, rightVal)
}

func NewTree[L, R, T any](
	reduce func(L, R) T,
	left func() L,
	right func() R,
) *Tree[L, R, T] {
	return &Tree[L, R, T]{
		reduce: reduce,
		left:   left,
		right:  right,
	}
}
