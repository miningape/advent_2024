package day18_problem1

import (
	"advent2024/day18"
	"advent2024/util"
	"fmt"
)

type Day18Solution1 struct {}

const GRID_X, GRID_Y = 71, 71
const CORRUPTED_BYTES = 1024

func (Day18Solution1) Solve(path string) {
	file := util.ReadFile(path)
	corruptions := day18.ParseInput(file)

	grid := util.GridInit(GRID_X, GRID_Y, false)
	day18.ApplyCorruption(corruptions[:CORRUPTED_BYTES], grid)

	shortestPath, found := grid.ShortestPath(
		util.Vector{X: 0, Y: 0}, 
		util.Vector{X: GRID_X - 1, Y: GRID_Y - 1}, 
		func(location util.Vector) map[util.Vector]int {
			m := make(map[util.Vector]int)

			for _, cardinal := range util.Cardinals() {
				neighbor := location.Add(cardinal)
				if isWall, found := grid.At(neighbor); found && !isWall {
					m[neighbor] = 1
				}
			}
			
			return m
		})

	if !found {
		panic("Could not find a path.")
	}

	fmt.Println(day18.ToString(grid, shortestPath))
	fmt.Println("length:", len(shortestPath) - 1)
}