package util

type Counter[T comparable] map[T]int;

func CounterOf[T comparable](initial... T) Counter[T] {
	counter := make(Counter[T])

	for _, elem := range initial {
		counter.Add(elem)
	}

	return counter
}

func (c Counter[T]) IsEmpty() bool {
	for _, count := range c {
		if count > 0 {
			return false
		}
	}

	return true
}

func (c Counter[T]) Has(n T) bool {
	v, found := c[n]

	return found && v > 0
}

func (c Counter[T]) Add(n T) {
	v, found := c[n]

	if !found {
		c[n] = 1
		return
	}

	c[n] = v + 1
}

func (c Counter[T]) AddTimes(n T, count int) {
	v, found := c[n]

	if !found {
		c[n] = count
		return
	}

	c[n] = v + count
}

func (c Counter[T]) Remove(n T) {
	if !c.Has(n) {
		panic("Tried to remove number that doesn't exist from counter")
	}

	if c[n] == 1 {
		delete(c, n)
	} else {
		c[n] = c[n] - 1
	}
}

func (c Counter[T]) Collect(f func(T, int, int) int, sum int) int {
	for num, count := range c {
		sum = f(num, count, sum)
	}

	return sum
}