package util

import "slices"

type GridV[T any] map[Vector]T

func GridOfRaw(m []string) GridV[rune] {
	grid := make(GridV[rune])

	for y, line := range m {
		for x, c := range line {
			grid[Vector{x, y}] = c
		}
	}

	return grid
}

type Grid[T any] [][]T

func GridOfString[T any](m []string, ma func(rune, Vector) T) Grid[T] {
	grid := make(Grid[T], 0, len(m))

	for y, line := range m {
			grid = append(grid, make([]T, 0, len(line)))

			for x, c := range line {
				grid[y] = append(grid[y], ma(c, Vector{x, y}))
			}
		}

	return grid
}

func (grid Grid[T]) At(location Vector) (T, bool)  {
	if location.Y < 0 || location.Y >= len(grid) {
		var dflt T
		return dflt, false
	}

	line := grid[location.Y]
	if location.X < 0 || location.X >= len(line) {
		var dflt T
		return dflt, false
	}

	return line[location.X], true
}

func (grid Grid[T]) Set(location Vector, value T) bool  {
	if location.Y < 0 || location.Y >= len(grid) {
		return false
	}

	line := grid[location.Y]
	if location.X < 0 || location.X >= len(line) {
		return false
	}

	grid[location.Y][location.X] = value
	return true
}

func (grid *Grid[T]) DeepClone() Grid[T] {
	copy := make(Grid[T], 0, len(*grid))

	for _, line := range *grid {
		copy = append(copy, slices.Clone(line))
	}

	return copy
}
