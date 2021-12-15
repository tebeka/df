package df

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	require := require.New(t)

	size := 10
	c := make(Column[int], size)
	require.Equal(size, len(c))
}

func TestSum(t *testing.T) {
	require := require.New(t)

	ci := Column[int]{1, 2, 3}
	require.Equal(6, c.Sum())

	cs := Column[string]{"a", "b", "c"}
	require.Equal("abc", cs.Sum())
}
