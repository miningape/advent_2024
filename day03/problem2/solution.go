package day3_problem2

import (
	"advent2024/day03"
	"advent2024/util"
	"fmt"
)


type reader struct {
	*day03.Reader
}

func (r *reader) dont() {
	for !r.IsAtEnd() {
		if r.Scan("do()") {
			break
		}
	}
}

func (r *reader) read() int {
	sum := 0

	for !r.IsAtEnd() {
		if r.Speculate_scan("don't()") {
			r.dont()
		}

		if !r.Scan("mul(") {
			continue
		}

		left, ok := r.Scan_int()
		if !ok {
			continue
		}

		if !r.Scan(",") {
			continue
		}

		right, ok := r.Scan_int()
		if !ok {
			continue
		}

		if !r.Scan(")") {
			continue
		}

		sum += left * right
	}
	
	return sum
}


type Day3Solution2 struct {}

func (Day3Solution2) Solve(path string) {
	file := util.ReadFile(path)
	r := reader { &day03.Reader{ Source: file, Index: 0 } }

	s := r.read()

	fmt.Println(s)
}