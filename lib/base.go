package lib

import (
	"fmt"
)

type BaseSort interface {
	// any collection that imple sort() will sort this particular collection
	Sort()

	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type EnhanceBaseSort interface {
	BaseSort

	// defined a way to print this collection
	Show()

	IsSorted() bool
}

type IntSlice []int

func (c *IntSlice) Len() int {
	return len(*c)
}
func (c *IntSlice) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}
func (c *IntSlice) Less(i, j int) bool {
	return (*c)[i] < (*c)[j]
}
func (c *IntSlice) Show() {
	fmt.Printf("%v\n", *c)
}
