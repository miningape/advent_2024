package day11_problem1

import (
	"advent2024/day11"
	"advent2024/util"
	"fmt"
)

func blink(stones []int) []int {
	next := make([]int, 0)

	for _, stone := range stones {
		if stone == 0 {
			next = append(next, 1)
			continue
		}
		
		if digits := day11.CountDigits(stone); digits % 2 == 0 {
			n := day11.SplitInt(stone)
			left := n[:digits/2]
			right := n[digits/2:]

			next = append(next, day11.JoinInt(left))
			next = append(next, day11.JoinInt(right))
			continue
		}

		next = append(next, stone * 2024)
	}

	return next
}

type Day11Solution1 struct {}

func (Day11Solution1) Solve(path string) {
	file := util.ReadFile(path)
	stones := day11.ParseInput(file)

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Println(len(stones))
}