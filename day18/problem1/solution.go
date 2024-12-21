package day18_problem1

import (
	"advent2024/util"
	"fmt"
	"strconv"
	"strings"
)

func ParseInput(file string) []util.Vector {
	lines := strings.Split(file, "\n")
	corrupted := make([]util.Vector, len(lines))

	for i, line := range lines {
		l := strings.Split(line, ",")
		
		x, err := strconv.Atoi(l[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}

		corrupted[i] = util.Vector{
			X: x,
			Y: y,
		}
	}

	return corrupted
}

func toString(memory util.Grid[bool], path []util.Vector) string {
	p := make(map[util.Vector]string)
	for _, step := range path {
		p[step] = "O"
	}

	 s, err := util.GridToString(memory, map[bool]string {
		true: "#",
		false: ".",
	}, p)

	if err != nil {
		panic(err)
	}

	return s
}

func applyCorruption(corruptions []util.Vector, grid util.Grid[bool]) {
	for _, corruption := range corruptions {
		grid.Set(corruption, true)
	}
}

type Day18Solution1 struct {}

const GRID_X, GRID_Y = 7, 7
const CORRUPTED_BYTES = 12

func (Day18Solution1) Solve(path string) {
	file := util.ReadFile(path)
	corruptions := ParseInput(file)

	grid := util.GridInit(GRID_X, GRID_Y, false)
	applyCorruption(corruptions[:CORRUPTED_BYTES], grid)

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

	// util.SetOf(shortestPath...)

	// for _, corruptions := range corruptions[CORRUPTED_BYTES:] 

	fmt.Println(toString(grid, shortestPath))
	fmt.Println("length:", len(shortestPath) - 1)
}