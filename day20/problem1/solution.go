package day20_problem1

import (
	"advent2024/day20"
	"advent2024/util"
	"fmt"
)

func shortestPathLength(maze util.Grid[bool], start, end, removedWall util.Vector) int {
	path, found := maze.ShortestPath(start, end, func(v util.Vector) map[util.Vector]int {
		neighbors := make(map[util.Vector]int)

		for _, cardinal := range util.Cardinals() {
			neighbor := v.Add(cardinal)
			
			if isWall, found := maze.At(neighbor); found && (!isWall || neighbor == removedWall) {
				neighbors[neighbor] = 1
			}
		}

		return neighbors
	})

	if !found {
		panic("Could not reach the end")
	}

	return len(path)
}

func findCheats(maze util.Grid[bool], skippable []util.Vector, start, end util.Vector) util.Counter[int] {
	sc := util.CounterOf[int]()

	for _, wall := range skippable {
		length := shortestPathLength(maze, start, end, wall)

		sc.Add(length)
	}

	return sc
}

func findSkippableWalls(grid util.Grid[bool], track []util.Vector) []util.Vector {
	skippable := util.SetOf[util.Vector]()

	for _, step := range track {
		for _, cardinal := range util.Cardinals() {
			neighbor := step.Add(cardinal)
			isWall, found := grid.At(neighbor)

			if found && isWall {
				skippable.Add(neighbor)
			}
		}
	}

	return skippable.Slice()
}

type Day20Solution1 struct {}

func (Day20Solution1) Solve(path string) {
	file := util.ReadFile(path)
	maze, start, end, track := day20.ParseInput(file)

	skippable := findSkippableWalls(maze, track.Slice())
	cheats := findCheats(maze, skippable, start, end)

	total := 0
	for saved, count := range cheats {
		if len(track) - saved  >= 100 {
			total += count
		}
	}

	fmt.Println(total)
}