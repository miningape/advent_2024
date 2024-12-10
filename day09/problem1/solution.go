package day9_problem1

import (
	"advent2024/day09"
	"advent2024/util"
	"fmt"
)

func writeFile(memory []int, diskMap []int, file int) []int {
	fileId := day09.GetFileId(file)
	fileSize := day09.GetFileSize(file, diskMap)

	next := memory

	for i := 0; i < fileSize; i++ {
		next = append(next, fileId)
	}

	return next
}

func move(to []int, from []int, size int) ([]int, []int) {
	for i := 0; i < size && i < len(from); i++ {
		to = append(to, from[i])
	}

	if size > len(from) {
		return to, make([]int, 0)
	}

	if size < 0 {
		return to, from
	}

	return to, from[size:]
}

func moveBack(to []int, from []int, size int) ([]int, []int) {
	for i := len(from) - 1; i >= len(from) - size && i >= 0; i-- {
		to = append(to, from[i])
	}

	if size > len(from) {
		return to, make([]int, 0)
	}

	return to, from[0:len(from) - size]
}

func compact(diskMap []int) []int {
	filesystem := make([]int, 0)
	buffer := make([]int, 0)

	last := 0
	for i, j := 0, len(diskMap) - 1; i <= j; i, j = i + 1, j - 1 {
		last = i
		// fmt.Println(i, j, filesystem, buffer)
		if i == j {
			// i = j is always even - is a file
			filesystem = writeFile(filesystem, diskMap, i)
			continue
		}

		if day09.IsFile(i) {
			filesystem = writeFile(filesystem, diskMap, i)

			if day09.IsFreeSpace(j) {
				continue
			}

			buffer = writeFile(buffer, diskMap, j)
		} else {
			fileSize := day09.GetFileSize(i, diskMap)

			if len(buffer) > 0 {
				newFileSize := fileSize - len(buffer)
				filesystem, buffer = move(filesystem, buffer, fileSize)

				fileSize = newFileSize
			}

			if day09.IsFreeSpace(j) {
				j--
			}

			temp := make([]int, 0, day09.GetFileSize(j, diskMap))
			temp = writeFile(temp, diskMap, j)
			for len(temp) < fileSize {
				j--

				if day09.IsFile(j) {
					temp = writeFile(temp, diskMap, j)
				}
			}

			filesystem, temp = move(filesystem, temp, fileSize)
			buffer, _ = move(buffer, temp, len(temp))
		}
	}

	for i := last + 1; len(buffer) > 0; i++ {
		fileSize := day09.GetFileSize(i, diskMap)
		if day09.IsFile(i) {
			filesystem, buffer = moveBack(filesystem, buffer, fileSize)
		} else {
			filesystem, buffer = move(filesystem, buffer, fileSize)
		}
	}

	filesystem, _ = move(filesystem, buffer, len(buffer))
	

	return filesystem
}


type Day9Solution1 struct {}

func (Day9Solution1) Solve(path string) {
	file := util.ReadFile(path)
	diskMap := day09.ParseInput(file)

	filesystem := compact(diskMap)
	checksum := day09.Checksum(filesystem)

	fmt.Println("Answer:", checksum)
}