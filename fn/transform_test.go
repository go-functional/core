package fn

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompose(t *testing.T) {
	r := require.New(t)
	composedFn := Compose(
		func(i int) string { return strconv.Itoa(i) },
		func(s string) string { return fmt.Sprintf("str-%s", s) },
	)
	r.Equal("str-123", composedFn(123))

}
