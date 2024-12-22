package day19

import (
	"advent2024/util"
	"strings"
)

func parseTowels(towels string) []string {
	return strings.Split(towels, ", ")
}

func parseDesigns(designs string) []string {
	return strings.Split(designs, "\n")
}

func ParseInput(file string) ([]string, []string) {
	f := strings.Split(file, "\n\n")
	return parseTowels(f[0]), parseDesigns(f[1])
}

func TowelsInDesign(towels []string, design string) map[string][]int {
	found := make(map[string][]int)

	for _, towel := range towels {
		found[towel] = util.FindSubStrings(design, towel)
	}

	for towel, locations := range found {
		if len(locations) == 0 {
			delete(found, towel)
		}
	}

	return found
}