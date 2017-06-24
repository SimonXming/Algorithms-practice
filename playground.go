package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	x, a := a[0], a[1:]
	fmt.Println(a, x)
}
