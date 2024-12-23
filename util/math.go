package util

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func Max(left int, right int) int {
	if left > right {
		return left
	}

	return right
}

func Min(left int, right int) int {
	if left > right {
		return right
	}

	return left
}

func SumList(s []int) uint64 {
	var sum uint64 = 0

	for _, e := range s {
		sum += uint64(e)
	}

	return sum
}

func Reverse(s []int) []int {
	size := len(s)
	opposite := make([]int, size)

	for i, e := range s {
		opposite[size - 1 - i] = e
	}

	return opposite
}

func Pow(base int, exponent int) int {
	pow := 1

	for i := 0; i < exponent; i++ {
		pow *= base
	}

	return pow
}
