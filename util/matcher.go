package util

import (
	"fmt"
	"strconv"
	"unicode"
)

type Matcher struct {
	Source string
	Index  int
}

func (t *Matcher) IsAtEnd() bool {
	return !(t.Index < len(t.Source))
}

func (t *Matcher) GetCurrent() rune {
	return rune(t.Source[t.Index])
}

func (t *Matcher) Advance() rune {
	c := t.GetCurrent()
	t.Index++
	return c
}

func (t *Matcher) Scan(s string) bool {
	for _, r := range s {
		c := t.Advance()

		if c != r {
			return false
		}
	}

	return true
}

func (t *Matcher) Scan_int() (int, bool) {
	s := ""
	c := t.GetCurrent()

	for unicode.IsDigit(c) && !t.IsAtEnd() {
		s += string(c)

		t.Advance()
		if !t.IsAtEnd() {
			c = t.GetCurrent()
		}
	}

	if len(s) == 0 {
		return 0, false
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return 0, false
	}

	return n, true
}

func (t *Matcher) Speculate_scan(s string) bool {
	if len(s) == 0 {
		panic("Cannot speculate_scan with an empty string")
	}

	first := rune(s[0])
	if t.GetCurrent() != first {
		return false
	}

	return t.Scan(s)
}

func (t *Matcher) Match(s string) {
	if !t.Scan(s) {
		panic(fmt.Sprint("Could not match \"", s, "\" in \n\"", t.Source[t.Index-10:t.Index+10], "\"\n at \"", t.Index))
	}
}

func (t *Matcher) Match_int() int {
	i, found := t.Scan_int()
	if !found{
		panic(fmt.Sprint("Could not match int in", t.Source[t.Index-10:t.Index+10], "at", t.Index))
	}

	return i
}