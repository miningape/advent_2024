package day11_problem2

import (
	"advent2024/day11"
	"advent2024/util"
	"fmt"
)

func blink(stone int, cache map[int][]int) []int {
	next, found := cache[stone]
	if found {
		return next
	}
		
	if digits := day11.CountDigits(stone); digits % 2 == 0 {
		n := day11.SplitInt(stone)
		left := n[:digits/2]
		right := n[digits/2:]

		next = append(next, day11.JoinInt(left))
		next = append(next, day11.JoinInt(right))
	} else {
		next = append(next, stone * 2024)
	}

	cache[stone] = next
	return next
}

func howManyStonesAfter(times int, stones util.Counter[int]) int {
	cache := map[int][]int { 0: { 1 } }

	for i := 0; i < times; i++ {
		next := util.CounterOf[int]()

		for stone, count := range stones {
			stones := blink(stone, cache)
			for _, stone := range stones {
				next.AddTimes(stone, count)
			}
		}

		stones = next
	}

	return stones.Collect(func(stone, amount, total int) int {
		return total + amount
	}, 0)
}

type Day11Solution2 struct {}

func (Day11Solution2) Solve(path string) {
	file := util.ReadFile(path)
	stones := util.CounterOf(day11.ParseInput(file)...)

	total := howManyStonesAfter(75, stones)

	fmt.Println(total)
}