package day12_problem1

import (
	"advent2024/day12"
	"advent2024/util"
	"fmt"
	"strings"
)

func calculatePerimeter(edges map[util.Vector]int) int {
	perimeter := 0

	for _, edgeCount := range edges {
		perimeter += edgeCount
	}

	return perimeter
}

type Day12Solution1 struct {}

func (Day12Solution1) Solve(path string) {
	file := util.ReadFile(path)
	grid := util.GridOfRaw(strings.Split(file, "\n"))

	plots := day12.FindAllPlots(grid)

	price := 0
	for _, plot := range plots {
		perimeter := calculatePerimeter(plot.Edges)
		price += plot.Area * perimeter
	}

	fmt.Println(price)
}