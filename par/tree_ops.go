package par

// AddLeft adds left to the left side of root
// and returns the new tree with the left
// side attached. If a left side already existed on root,
// it is replaced.
func AddLeft[L, L2, R, R2, T any](
	root Tree[L, R, T],
	left Tree[L2, R2, L],
) Tree[L, R, T] {
	return AddLeftFunc[L, L2, R, R2](
		root,
		left.Exec,
	)
}

// AddLeftFunc adds left to the left side of root and returns
// the new resultant tree.
// If a left side already existed on root, it is replaced.
//
// This function is generally not used directly because
// it does not easily allow you to specify a sub-tree
// of the left side. Consider using AddLeft instead
func AddLeftFunc[L, L2, R, R2, T any](
	root Tree[L, R, T],
	left func() L,
) Tree[L, R, T] {
	root.left = left
	return root
}

// AddRight adds right to the right side of root
// and returns the new tree with the right
// side attached. If a right side already existed on root,
// it is replaced.
func AddRight[L, L2, R, R2, T any](
	root Tree[L, R, T],
	right Tree[L2, R2, R],
) Tree[L, R, T] {
	return AddRightFunc[L, L2, R, R2](root, right.Exec)
}

// AddRightFunc adds right to the right side of root and
// returns the new resultant tree.
// If a right side already existed on root, it is replaced.
//
// This function is generally not used directly because
// it does not easily allow you to specify a sub-tree
// of the right side. Consider using AddRight instead
func AddRightFunc[L, L2, R, R2, T any](
	root Tree[L, R, T],
	right func() R,
) Tree[L, R, T] {
	root.right = right
	return root
}
