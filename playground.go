package main

import "fmt"

// Parent Object.
type Object struct{}

func (b Object) String() string {
	return "I am an Object"
}

type ObjectOne struct {
	Object
}

type ObjectTwo struct {
	Object
}

func (b ObjectTwo) String() string {
	return "I am an Object Two"
}

func main() {
	o := Object{}
	fmt.Println(o) // I am an Object

	o1 := ObjectOne{}
	fmt.Println(o1) // I am an Object

	o2 := ObjectTwo{}
	fmt.Println(o2) // I am an Object Two
}
