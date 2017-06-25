package main

import "fmt"

func merge(a []int, lo int, mid int, hi int) {
	/*
		原位归并
		a[lo..mid] 和 a[mid+1..hi] 归并
	*/
	// i 和 j 是分割为两个子数组的标志位
	// 首先初始化为起始位置标志位，一个在开头一个在中间
	i := lo
	j := mid + 1

	aux := make([]int, 0)

	for _, item := range a {
		aux = append(aux, item)
	}
	fmt.Printf("%v\n", aux)
	// 从 lo 遍历到 hi
	for k := lo; k <= hi; k++ {
		if i > mid {
			// 如果 i 大于 mid 说明前一个数组已经排序完毕了
			// 全部从第二个数组取值，并按顺序放在 a 中
			a[k] = aux[j]
			j++
		} else if j > hi {
			// 如果 j 大于 hi 说明后一个数组已经排序完毕了
			// 全部从第二个数组取值，并按顺序放在 a 中
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			// 取 i, j 标志位中更小的值按顺序放在 a 里
			// j < i; aux[j] 放入 a 中
			a[k] = aux[j]
			j++
		} else {
			// i < j; aux[i] 放入 a 中
			a[k] = aux[i]
			i++
		}
	}

	fmt.Printf("%v\n", a)
}

func main() {
	a := []int{1, 3, 6, 7, 8, 0, 2, 4, 5}
	merge(a, 0, 4, 8)
}
