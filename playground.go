package main

import (
	"fmt"
)

type testInt func() bool // 声明了一个函数类型

func main() {
	// x := base_struct_one.Item{
	// 	Data: "data",
	// }
	// y := base_struct_one.Bag()
	// fmt.Printf(x.Data + "\n")
	// fmt.Printf(y)
	// fmt.Printf("hello, world, %v\n", base_struct_one.Sqrt(2))
	fmt.Printf("hello, world %v\n", "!")
	throwsPanic(some_panic)
}

func some_panic() bool {
	panic("no value for $USER")
	return true
}

func throwsPanic(f testInt) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
