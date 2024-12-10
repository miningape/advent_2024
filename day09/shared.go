package day09

import (
	"strconv"
	"strings"
)

func ParseInput(file string) []int {
	list := strings.Split(file, "")
	next := make([]int, 0, len(list))

	for _, elem := range list {
		num, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}

		next = append(next, num)
	}

	return next
}

func Checksum(filesystem []int) int {
	sum := 0

	for i, block := range filesystem {
		sum += i * block
	}

	return sum
}

func IsFile(n int) bool {
	return n % 2 == 0
}

func IsFreeSpace(n int) bool {
	return n % 2 == 1
}

func GetFileId(n int) int {
	return n / 2
}

func GetFileSize(n int, diskMap []int) int {
	return diskMap[n]
}
