package day03

import (
	"fmt"
	"strconv"
	"unicode"
)

type Reader struct {
	Source string
	Index int
}

func (t *Reader) GetCurrent() rune {
	return rune(t.Source[t.Index])
}

func (t *Reader) Advance() rune {
	c := t.GetCurrent()
	t.Index++
	return c
}

func (t *Reader) Scan(s string) bool {
	for _, r := range s {
		c := t.Advance()
		
		if c != r {
			return false
		}
	}

	return true
}

func (t *Reader) Scan_int() (int, bool) {
	c := t.GetCurrent()
	s := ""
	for unicode.IsDigit(c) && len(s) < 3 {
		s += string(c)
		t.Advance()
		c = t.GetCurrent()
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

func (t *Reader) Speculate_scan(s string) bool {
	if len(s) == 0 {
		panic("Cannot speculate_scan with an empty string")	
	}

	first := rune(s[0])
	if t.GetCurrent() != first {
		return false
	}

	return t.Scan(s)
}