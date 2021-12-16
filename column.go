package df

import (
	"constraints"
	"fmt"
	"sort"
)


type Value interface {
	// TODO: time.Time
	constraints.Ordered
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

func (c Column[T]) edge(name string, fn func(a, b T) bool) (T, error){
	if len(c) == 0 {
		var z T
		return z, fmt.Errorf("%s of empty column", name)
	}

	m := c[0]
	for _, v := range c[1:] {
		if fn(v, m) {
			m = v
		}
	}
	return m, nil
}

func (c Column[T]) Max() (T, error) {
	return c.edge("Max", func(a, b T) bool { return a > b})
}

func (c Column[T]) Min() (T, error) {
	return c.edge("Min", func(a, b T) bool { return a < b})
}
