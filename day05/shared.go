package day05

import (
	"advent2024/util"
	"strconv"
	"strings"
)


func getOrderingRules(file string) map[int]util.Set[int] {
	rules := make(map[int]util.Set[int])

	for _, rule := range strings.Split(file, "\n") {
		r := strings.Split(rule, "|")
		before, err := strconv.Atoi(r[0])
		if err != nil {
			panic(err)
		}

		after, err := strconv.Atoi(r[1])
		if err != nil {
			panic(err)
		}

		_, found := rules[before]
		if !found {
			rules[before] = util.SetOf[int]()
		}
		
		rules[before].Add(after)
	}

	return rules
}

func getPageUpdates(file string) [][]int {
	lines := strings.Split(file, "\n")
	updates := make([][]int, 0, len(lines))

	for _, line := range lines {
		page_strings := strings.Split(line, ",")
		pages := make([]int, 0, len(page_strings))

		for _, s := range page_strings {
			page, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}

			pages = append(pages, page)
		}

		updates = append(updates, pages)
	}

	return updates
}

func ParseInput(file string) (map[int]util.Set[int], [][]int) {
	parts := strings.Split(file, "\n\n")

	rules := getOrderingRules(parts[0])
	updates := getPageUpdates(parts[1])

	return rules, updates
}

func FindMiddle(l []int) int {
	if len(l) % 2 == 0 {
		panic("Cannot find the middle of an even list")
	}

	middle := (len(l) - 1) / 2
	return l[middle]
}

func IsCorrectlyOrdered(pages []int, rules map[int]util.Set[int]) bool {
	seen := util.SetOf[int]()

	for _, page := range pages {
		mustBeBefore := rules[page]
		if len(mustBeBefore.Intersection(seen)) > 0 {
			return false
		}

		seen.Add(page)
	}
	
	return true
}