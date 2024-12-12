package day11

import (
	"math"
	"strconv"
	"strings"
)

func ParseInput(file string) []int {
	nums := strings.Split(file, " ")
	stones := make([]int, 0, len(nums))

	for _, num := range nums {
		stone, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		stones = append(stones, stone)
	}

	return stones
}

func CountDigits(num int) int {
	return int(math.Floor(math.Log10(float64(num)))) + 1
}

func SplitInt(num int) []int {
	slc := []int{}
	for num > 0 {
		slc = append(slc, num % 10)
		num /= 10
	}

	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}

	return slc
}

func JoinInt(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum = sum * 10
		sum += num
	}

	return sum
}