package day12

import "advent2024/util"

type CropLocation struct {
	Crop rune
	Position util.Vector
}

type CropPlot struct {
	Edges map[util.Vector]int
	Area int
}

func exploreFromCropLocation(grid util.Grid[rune], location CropLocation, explored util.Set[util.Vector]) CropPlot {
	cardinals := util.Cardinals()
	stack := util.StackOf(location.Position)
	plot := CropPlot{
		make(map[util.Vector]int),
		0,
	}

	for !stack.IsEmpty() {
		var position util.Vector
		position, stack = stack.Pop()

		if explored.Contains(position) {
			continue
		}

		explored.Add(position)
		plot.Area++

		for _, direction := range cardinals {
			next := position.Add(direction)
			crop, found := grid[next]
			if !found {
				plot.Edges[position]++
				continue
			}

			if crop == location.Crop {
				stack = stack.Push(next)
			} else {
				plot.Edges[position]++
			}
		}
	}

	return plot
}

func FindAllPlots(grid util.Grid[rune]) map[CropLocation]CropPlot {
	explored := util.SetOf[util.Vector]()
	plots := make(map[CropLocation]CropPlot)

	for position, crop := range grid {
		if explored.Contains(position) {
			continue
		}

		location := CropLocation{ crop, position }
		plot := exploreFromCropLocation(grid, location, explored)
		plots[location] = plot
	}

	return plots
}