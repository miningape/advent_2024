package day17_problem1

import (
	"advent2024/day17"
	"advent2024/util"
	"fmt"
	"strconv"
)

func toString(output []int) string {
	s := ""

	for _, result := range output {
		if len(s) > 0 {
			s += ","
		}

		s += strconv.Itoa(result)
	}

	return s
}

type Day17Solution1 struct{}

func (Day17Solution1) Solve(path string) {
	file := util.ReadFile(path)
	computer := day17.ParseInput(file)

	for !computer.IsHalted() {
		computer = computer.Cycle()
	}

	fmt.Println(toString(computer.Output))
}