package par

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	left := NewTree(
		func(l string, r int) string {
			return fmt.Sprintf("l=%s,r=%d", l, r)
		},
		func() string { return "left" },
		func() int { return 3 },
	)
	right := NewTree(
		func(l string, r string) string {
			return fmt.Sprintf("l=%s,r=%s", l, r)
		},
		func() string { return "right->left" },
		func() string { return "right->right" },
	)
	root := NewTree(
		func(l string, r string) string {
			return fmt.Sprintf("l=%s,r=%s", l, r)
		},
		left.Exec,
		right.Exec,
	)
	ret := root.Exec()
	t.Logf(ret)

}
