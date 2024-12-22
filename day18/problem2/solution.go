package day18_problem2

import (
	"advent2024/day18"
	"advent2024/util"
	"fmt"
)

func getShortestPath(memory util.Grid[bool]) ([]util.Vector, bool) {
	return memory.ShortestPath(
		util.Vector{X: 0, Y: 0}, 
		util.Vector{X: GRID_X - 1, Y: GRID_Y - 1}, 
		func(location util.Vector) map[util.Vector]int {
			m := make(map[util.Vector]int)

			for _, cardinal := range util.Cardinals() {
				neighbor := location.Add(cardinal)
				if isWall, found := memory.At(neighbor); found && !isWall {
					m[neighbor] = 1
				}
			}
			
			return m
	})
}

func corruptUntilUnsolveable(memory util.Grid[bool], corruptions []util.Vector) util.Vector {
	shortestPath, found := getShortestPath(memory)
	if !found {
		panic("Could not find a path on the first check. Must be solveable to be made unsolveable.")
	}

	steps := util.SetOf(shortestPath...)
	for i, corruption := range corruptions {
		memory.Set(corruption, true)
		
		if !steps.Contains(corruption) {
			continue
		}
		
		shortestPath, found := getShortestPath(memory)

		if !found {
			return corruptions[i]
		}

		steps = util.SetOf(shortestPath...)
	}

	panic("Could not corrupt the memory to be untraversable.")
}

type Day18Solution2 struct {}

const GRID_X, GRID_Y = 71, 71
const CORRUPTED_BYTES = 1024

func (Day18Solution2) Solve(path string) {
	file := util.ReadFile(path)
	corruptions := day18.ParseInput(file)

	memory := util.GridInit(GRID_X, GRID_Y, false)

	// We know that the map must be solveable since the last question asked us to find that path
	day18.ApplyCorruption(corruptions[:CORRUPTED_BYTES], memory)
	finalCorruption := corruptUntilUnsolveable(memory, corruptions[CORRUPTED_BYTES:])
	
	fmt.Printf("%d,%d\n", finalCorruption.X, finalCorruption.Y)
}