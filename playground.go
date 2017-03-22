package main

import (
	"fmt"
	"github.com/SimonXming/Algorithms-practice/lib"
	"sort"
)

type testInt func() bool // 声明了一个函数类型

func main() {
	var y lib.IntSlice
	x := lib.IntSlice{1, 2, 3, 4, 5}
	sort.Sort(x)

	fmt.Println(sort.Search(len(x), func(i int) bool { return x[i] >= 1 }))
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
