package day12_problem2

import (
	"advent2024/day12"
	"advent2024/util"
	"fmt"
	"math"
	"strings"
)

func minimalContainingBox(plot day12.CropPlot) (util.Vector, util.Vector) {
	max := util.Vector{ X: -1, Y: -1 }
	min := util.Vector{ X: math.MaxInt, Y: math.MaxInt }

	for edge := range plot.Edges {
		if edge.X > max.X {
			max.X = edge.X
		}

		if edge.Y > max.Y {
			max.Y = edge.Y
		}

		if edge.X < min.X  {
			min.X = edge.X
		}

		if edge.Y < min.Y {
			min.Y = edge.Y
		}
	}

	return max, min
}

func shadow(max, min, direction util.Vector) util.Vector {
	if direction.Collapse() < 0 {
		return max.VectorMul(direction.Opposite())
	}

	return min.VectorMul(direction)
}

// Finds the starting position so when we add direction and perpendicular we traverse the area we're interested in
// When a component (x, y) is negative we are walking backwards and therefore start from the max and travel to the min
func startPos(max, min, direction, perpendicular util.Vector) util.Vector {
	outer := shadow(max, min, direction)
	inner := shadow(max, min, perpendicular)

	return outer.Add(inner)
}

func calculateSidesOn(grid map[util.Vector]rune, location day12.CropLocation, plot day12.CropPlot, max, min, direction util.Vector) int {
	sides := 0
	bounds := max.Sub(min)
	previousBoundaries := util.SetOf[int]()
	perpendicular := direction.RotateOrigin90()
	start := startPos(max, min, direction, perpendicular)
	
	for outer := 0; outer <= direction.VectorMul(bounds).ManhattanOrigin(); outer++ {
		boundaries := util.SetOf[int]()
		isInside := false

		for inner := 0; inner <= perpendicular.VectorMul(bounds).ManhattanOrigin(); inner++ {
			current := start.Add(direction.Mul(outer)).Add(perpendicular.Mul(inner))

			_, found := plot.Edges[current]
			if found {
				isInside = true
			} else {
				if isInside {
					// Handle the case we are inside the crop (but no edge so we are still inside)
					crop, found := grid[current]
					if !found || location.Crop != crop {
						isInside = false
						boundaries.Add(inner - 1)
					}
				}
			}
		}

		if isInside {
			boundaries.Add(perpendicular.VectorMul(bounds).ManhattanOrigin())
		}

		endedFences := previousBoundaries.Not(boundaries)
		sides += len(endedFences)
		previousBoundaries = boundaries
	}

	sides += len(previousBoundaries)
	return sides
}

type Day12Solution2 struct {}

func (Day12Solution2) Solve(path string) {
	file := util.ReadFile(path)
	grid := util.GridOfRaw(strings.Split(file, "\n"))

	plots := day12.FindAllPlots(grid)

	price := 0
	for location, plot := range plots {
		sides := 0
		max, min := minimalContainingBox(plot)

		for _, cardinal := range util.Cardinals() {
			sides += calculateSidesOn(grid, location, plot, max, min, cardinal)
		}
		
		price += plot.Area * sides
	}

	fmt.Println(price)
}