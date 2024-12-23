package day20_problem2

import (
	"advent2024/day20"
	"advent2024/util"
	"fmt"
)

func nonCheatedPath(parentOf map[util.Vector]util.Vector, end util.Vector) map[util.Vector]int {
	step := 0
	path := make(map[util.Vector]int)

	parent, found := parentOf[end]
	for found {
		path[parent] = step
		step++

		parent, found = parentOf[parent]
	}

	return path
}

func useableTeleports(grid util.Grid[bool], track []util.Vector, pathOrder map[util.Vector]int) map[util.Vector][]util.Vector {
	teleportsFrom := make(map[util.Vector][]util.Vector)

	circle := util.ManhattanCircle(20)

	for _, to := range track {
		for offset := range circle {
			from := to.Add(offset)
			isWall, found := grid.At(from)

			if found && !isWall {
				if pathOrder[to] < pathOrder[from] {
					continue
				}

				if _, found := teleportsFrom[to]; !found {
					teleportsFrom[to] = make([]util.Vector, 0)
				}

				teleportsFrom[to] = append(teleportsFrom[to], from)
			}
		}
	}

	return teleportsFrom
}

func shortestPaths(maze util.Grid[bool], from util.Vector) map[util.Vector]util.Vector {
	return maze.ShortestPathsFrom(from, func(v util.Vector) map[util.Vector]int {
		neighbors := make(map[util.Vector]int)

		for _, cardinal := range util.Cardinals() {
			neighbor := v.Add(cardinal)
			
			if isWall, found := maze.At(neighbor); found && !isWall {
				neighbors[neighbor] = 1
			}
		}

		return neighbors
	})
}

func changeInPathLengthWithTeleport(pathOrder map[util.Vector]int, teleportFrom, teleportTo util.Vector) int {
	return (pathOrder[teleportFrom] - pathOrder[teleportTo]) + teleportFrom.Sub(teleportTo).ManhattanOrigin()
}

type Day20Solution2 struct {}

func (Day20Solution2) Solve(path string) {
	file := util.ReadFile(path)
	maze, start, end, track := day20.ParseInput(file)

	parentOf := shortestPaths(maze, start)
	
	pathOrder := nonCheatedPath(parentOf, end)
	teleports := useableTeleports(maze, track.Slice(), pathOrder)

	total := 0
	for to, f := range teleports {
		for _, from := range f {
			saved := changeInPathLengthWithTeleport(pathOrder, from, to) 

			if -saved >= 100 {
				total++
			}
		}
	}

	fmt.Println(total)
}
