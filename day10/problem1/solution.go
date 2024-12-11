package day10_problem1

import (
	"advent2024/day10"
	"advent2024/util"
	"fmt"
)

func findTrailheadScore(topography map[util.Vector]int, trailHead util.Vector) int {
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

		height, found := topography[position]
		if !found {
			continue
		}

		if height == 9 {
			score++
		}

		for _, dir := range dirs {
			n := position.Add(dir)
			next, found := topography[n]
			
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

type Day10Solution1 struct {}

func (Day10Solution1) Solve(path string) {
	file := util.ReadFile(path)
	topography, trailheads := day10.ParseTopography(file)

	sum := 0
	for _, trailhead := range trailheads {
		sum += findTrailheadScore(topography, trailhead)
	}

	fmt.Println(sum)
}