package main

import (
	"github.com/SimonXming/Algorithms-practice/lib"
)

func newCollection() *lib.IntSlice {
	return &lib.IntSlice{1, 6, 3}
}

func main() {
	l := newCollection()
	l.Show()
}
