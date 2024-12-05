package day5_problem2

import (
	"advent2024/util"
	"fmt"
	"sort"
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

func isOk(pages []int, rules map[int]util.Set[int]) bool {
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

func findMiddle(l []int) int {
	if len(l) % 2 == 0 {
		panic("Cannot find the middle of an even list")
	}

	middle := (len(l) - 1) / 2
	return l[middle]
}

type Day5Solution2 struct {}

func (Day5Solution2) Solve(path string) {
	file := strings.Split(util.ReadFile(path), "\n\n")

	rules := getOrderingRules(file[0])
	updates := getPageUpdates(file[1])

	sum := 0
	for _, update := range updates {
		if !isOk(update, rules) {
			sort.Slice(update, func(i, j int) bool {
				left := update[i]
				right := update[j]
				rule, found := rules[left]

				if found && rule.Contains(right) {
					return true
				}

				return false

			})

			sum += findMiddle(update)
		}
	}

	fmt.Println(sum)
}