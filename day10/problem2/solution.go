package day10_problem2

import (
	"advent2024/day10"
	"advent2024/util"
	"fmt"
)

func findPathsFromTrailhead(topography map[util.Vector]int, trailHead util.Vector) map[util.Vector][]util.Vector {
	children := make(map[util.Vector][]util.Vector)
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

		for _, dir := range dirs {
			n := position.Add(dir)
			next, found := topography[n]
			
			if !found {
				continue
			}

			if next == height + 1 {
				stack = stack.Push(n)

				_, found := children[position]
				if !found {
					children[position] = make([]util.Vector, 0)
				}

				children[position] = append(children[position], n)
			}
		}
	}
	
	return children
}

func trailheadRating(topography map[util.Vector]int, childrenOf map[util.Vector][]util.Vector, trailhead util.Vector) int {
	c, found := childrenOf[trailhead]
	if !found {
		if topography[trailhead] == 9 {
			return 1
		} else {
			return 0
		}
	}

	rating := 0

	for _, next := range c {
		rating += trailheadRating(topography, childrenOf, next)
	}

	return rating
}

type Day10Problem2 struct {}

func (Day10Problem2) Solve(path string) {
	file := util.ReadFile(path)
	topography, trailheads := day10.ParseTopography(file)

	sum := 0
	for _, trailhead := range trailheads {
		children := findPathsFromTrailhead(topography, trailhead)
		rating := trailheadRating(topography, children, trailhead)

		sum += rating
	}

	fmt.Println(sum)
}