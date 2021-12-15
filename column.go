package df

import (
	"constraints"
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
