package main

import (
	"fmt"
	"github.com/SimonXming/Algorithms-practice/base_struct_one"
)

func main() {
	x := base_struct_one.Item{
		Data: "data",
	}
	// y := base_struct_one.Bag()
	fmt.Printf(x.Data + "\n")
	// fmt.Printf(y)
	fmt.Printf("hello, world, %v\n", base_struct_one.Sqrt(2))
}
