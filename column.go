package df

import (
	"constraints"
	"fmt"
	"sort"
	"time"
)


type Value interface {
	// TODO: time.Time
	constraints.Ordered | time.Duration
}

type Column[T Value] []T

func (c Column[T]) Sum() T {
	var total T
	for _, v := range c {
		total += v
	}
	return total
}

func (c Column[T]) Sort() {
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})
}

func (c Column[T]) Max() (T, error) {
	if len(c) == 0 {
		var z T
		return z, fmt.Errorf("max of empty column")
	}
	m := c[0]
	for _, v := range c[1:] {
		if v > m {
			m = v
		}
	}
	return m, nil
}
