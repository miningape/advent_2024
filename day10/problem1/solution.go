package day10_problem1

import (
	"advent2024/day10"
	"advent2024/util"
	"fmt"
)

func findTrailheadScore(topograph map[util.Vector]int, trailHead util.Vector) int {
	score := 0
	dirs := day10.AdjacentDirections()
	stack := util.StackOf(trailHead)
	discovered := util.SetOf[util.Vector]()
	
	for !stack.IsEmpty() {
		var position util.Vector
		position, stack = stack.Pop()

		if discovered.Contains(position) {
			continue
		}

		discovered.Add(position)

		height, found := topograph[position]
		if !found {
			continue
		}

		if height == 9 {
			score++
		}

		for _, dir := range dirs {
			n := position.Add(dir)
			next, found := topograph[n]
			
			if !found {
				continue
			}

			if next == height + 1 {
				stack = stack.Push(n)
			}
		}
	}
	
	return score
}

type Day10Problem1 struct {}

func (Day10Problem1) Solve(path string) {
	file := util.ReadFile(path)
	topograph, trailheads := day10.ParseTopograph(file)

	sum := 0
	for _, trailhead := range trailheads {
		sum += findTrailheadScore(topograph, trailhead)
	}

	fmt.Println(sum)
}