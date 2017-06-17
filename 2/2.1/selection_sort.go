/*
选择排序
对于长度为 N 的数组，选择排序需要大约 N^2/2 次比较和 N 次交换。
*/
package main

import (
	"fmt"
)

type CharSlice []string

func (c *CharSlice) Len() int {
	return len(*c)
}
func (c *CharSlice) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}
func (c *CharSlice) Less(i, j int) bool {
	return (*c)[i] < (*c)[j]
}
func (c *CharSlice) Show() {
	fmt.Printf("%v\n", *c)
}

func (c *CharSlice) Sort() {
	length := c.Len()
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if c.Less(j, min) {
				println((*c)[j], "<", (*c)[min])
				min = j
			}
			c.Swap(i, min)
		}
	}
}

func newCollection() *CharSlice {
	return &CharSlice{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
}

func main() {
	s_slict := newCollection()
	s_slict.Sort()
	s_slict.Show()
}
