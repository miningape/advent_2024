package util

import (
	"errors"
	"fmt"
	"slices"
)

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

func GridInit[T any](x, y int, dflt T) Grid[T] {
	grid := make(Grid[T], 0, y)

	for Y := 0; Y < y; Y++ {
		row := make([]T, x)

		for X := 0; X < x; X++ {
			row[X] = dflt
		}

		grid = append(grid, row)
	}

	return grid
}

func GridOfString[T any](m []string, mapper func(rune, Vector) T) Grid[T] {
	grid := make(Grid[T], 0, len(m))

	for y, line := range m {
			grid = append(grid, make([]T, 0, len(line)))

			for x, c := range line {
				grid[y] = append(grid[y], mapper(c, Vector{x, y}))
			}
		}

	return grid
}

func GridToString[T comparable](grid Grid[T], mapping map[T]string, positional map[Vector]string) (string, error) {
	g := ""

	for y, line := range grid {
		for x, cell := range line {
			representation, found := mapping[cell]
			if !found {
				return "", errors.New(fmt.Sprint("No mapping provided for \"", cell, "\" provided: ", mapping))
			}

			positionalRepresentation, found := positional[Vector{x, y}]
			if !found {
				g += representation
			} else {
				g += positionalRepresentation
			}
		}

		g += "\n"
	}

	return g, nil
}

func (grid Grid[T]) IsInside(location Vector) bool {
	if location.X < 0 || location.Y < 0 || location.Y >= len(grid) {
		return false
	}

	line := grid[location.Y]
	return location.X < len(line)
}

func (grid Grid[T]) At(location Vector) (T, bool)  {
	if !grid.IsInside(location) {
		var dflt T
		return dflt, false
	}

	return grid[location.Y][location.X], true
}

func (grid Grid[T]) Set(location Vector, value T) bool  {
	if !grid.IsInside(location) {
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

func (grid *Grid[T]) ShortestPath(from, to Vector, neighbors func(Vector) map[Vector]int) ([]Vector, bool) {
	pq := PriorityQueueOf(from, 0)
	parent := make(map[Vector]Vector)

	for !pq.IsEmpty() {
		location, cost := pq.PullPriority()

		if location == to {
			break
		}

		for neighbor, neighborCost := range neighbors(location) {
			_, found := parent[neighbor]
			if found || neighbor == from { continue }

			pq.Insert(neighbor, cost + neighborCost)
			parent[neighbor] = location
		}
	}

	path := make([]Vector, 0)

	p := to
	found := true
	for found {
		path = append(path, p)
		p, found = parent[p]
	}

	_, found = parent[to]
	return path, found
}
