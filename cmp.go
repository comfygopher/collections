package coll

import "cmp"

type valuesCounter[V cmp.Ordered] struct {
	counter map[V]int
}

func newValuesCounter[V cmp.Ordered]() *valuesCounter[V] {
	return &valuesCounter[V]{
		counter: make(map[V]int),
	}
}

func (c *valuesCounter[V]) Count(v V) int {
	if _, ok := c.counter[v]; !ok {
		return 0
	}
	return c.counter[v]
}

func (c *valuesCounter[V]) Increment(v V) {
	if _, ok := c.counter[v]; !ok {
		c.counter[v] = 1
	} else {
		c.counter[v]++
	}
}

func (c *valuesCounter[V]) Decrement(v V) {
	count, exists := c.counter[v]
	if !exists {
		return
	}
	if count == 1 {
		delete(c.counter, v)
	} else {
		c.counter[v]--
	}
}
