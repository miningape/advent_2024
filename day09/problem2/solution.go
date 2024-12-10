package day9_problem2

import (
	"advent2024/day09"
	"advent2024/util"
	"fmt"
	"time"
)

func getFileAddresses(diskMap []int) map[int]int {
	addresses := make(map[int]int)

	pointer := 0
	for i, size := range diskMap {
		if day09.IsFile(i) {
			addresses[day09.GetFileId(i)] = pointer
		}

		pointer += size
	}

	return addresses
}

func createFileSystem(diskMap []int) []int {
	filesystem := make([]int, 0)

	for i, size := range diskMap {
		for j := 0; j < size; j++ {
			if day09.IsFile(i) {
				filesystem = append(filesystem, day09.GetFileId(i))
			} else {
				filesystem = append(filesystem, 0)
			}
		}
	}

	return filesystem
}

func move(filesystem, diskMap []int, addresses map[int]int, location, file int) []int {
	fileId := day09.GetFileId(file)
	fileSize := day09.GetFileSize(file, diskMap)

	address := addresses[fileId]

	for i := 0; i < fileSize; i++ {
		filesystem[location + i] = filesystem[address + i]
		filesystem[address + i] = 0
	}

	return filesystem
}

func findNextFreeBlock(filesystem []int, skip int, until int) ([2]int, int, bool) {
	for skip < until && filesystem[skip] != 0  {
		skip++
	}

	if skip == until {
		return [2]int {}, skip, false
	}

	start := skip
	for skip < until && filesystem[skip] == 0 {
		skip++
	}

	return [2]int { start, skip - start }, skip, true
}

func findFreeBlocks(filesystem []int, skip int, until int) [][2]int {
	blocks := make([][2]int, 0)

	for skip < until {
		block, s, found := findNextFreeBlock(filesystem, skip, until)

		skip = s

		if found {
			blocks = append(blocks, block)
		}
	}

	return blocks
}

func shuffle(diskMap []int) []int {
	offset := (len(diskMap) - 1) % 2
	addresses := getFileAddresses(diskMap)
	filesystem := createFileSystem(diskMap)

	// fmt.Println(offset, len(filesystem), addresses, blockAddresses, filesystem)

	for file := len(diskMap) - 1 - offset; file >= 0; file -= 2 {
		fileSize := day09.GetFileSize(file, diskMap)
		address := addresses[day09.GetFileId(file)]
		// fmt.Println("file", file, day09.GetFileId(file), "size", fileSize, "address", address)

		blocks := findFreeBlocks(filesystem, day09.GetFileSize(0, diskMap), address)
		// fmt.Println(" ", blocks)
		for _, block := range blocks {
			// fmt.Println("- block", block)
			blockSize := block[1]

			if blockSize >= fileSize {
				// fmt.Println("- - found")
				location := block[0]
				filesystem = move(filesystem, diskMap, addresses, location, file)
				// fmt.Println("- -", filesystem)
				break
			}
		}
	}

	return filesystem
}

type Day9Solution2 struct {}

func (Day9Solution2) Solve(path string) {
	defer util.MeasureRuntime(time.Now())
	file := util.ReadFile(path)
	diskMap := day09.ParseInput(file)

	filesystem := shuffle(diskMap)
	checksum := day09.Checksum(filesystem)

	// fmt.Println(diskMap)
	// fmt.Println(filesystem)
	fmt.Println("Answer:", checksum)
}