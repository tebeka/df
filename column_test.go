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
	require.Equal(6, ci.Sum())

	cs := Column[string]{"a", "b", "c"}
	require.Equal("abc", cs.Sum())
}

func TestSort(t *testing.T) {
	require := require.New(t)

	cf := Column[float64]{2.2, 1.1, 3.3}
	cf.Sort()
	require.Equal(Column[float64]{1.1, 2.2, 3.3}, cf)

	cs := Column[string]{"B", "a", "A"}
	cs.Sort()
	require.Equal(Column[string]{"A", "B", "a"}, cs)
}

func TestMax(t *testing.T) {
	require := require.New(t)

	var ci Column[int]
	_, err := ci.Max()
	require.Error(err)

	ci = Column[int]{7}
	v, err := ci.Max()
	require.NoError(err)
	require.Equal(ci[0], v)
}
