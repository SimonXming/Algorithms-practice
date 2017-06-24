/*
希尔排序
基于插入排序的快速的排序算法，插入只会交换相邻的元素。
希尔排序为了加快插入排序：
1. 可以交换不相邻的元素
2. 对数组局部进行排序
3. 最终使用插入排序将局部有序的数组排序

希尔排序的思想是使数组中任意相隔为 h 的元素都是有序的，称为 h 有序数组。
换句话说，一个 h 有序数就是 h 个相互独立的有序数组编织在一起组成的数组。
在排序时，如果 h 很大，我们就能讲元素移动到很远的地方，为实现更小的 h 有序创造方便。

h = 4
L  E  E  A  M  H  L  E  P  S  O  L  T  S  X  R
L --------- M --------- P --------- T
   E --------- H --------- S --------- S
      E --------- L --------- O --------- X
         A --------- E --------- L --------- R

实现希尔排序的一种方法是对于每个 h，用插入排序将 h 个子数组独立的排序。
但因为子数组是相互独立的，一个更简单的方法是在 h 个子数组中将每个元素交换到比它大的元素前去。
只需要在插入排序的代码中将移动元素的距离由 1 改为 h 即可。

希尔排序更高效的原因是它权衡了子数组的规模和有序性。排序之前，各个子数组都很短，排序后子数组都是部分有序的，这两种情况都很适合插入排序

透彻理解希尔排序的性能至今仍是一项挑战。实际上，它是我们唯一无法准确描述其对于乱序的数组的性能特征的排序方式。
*/

package main

import (
	"bufio"
	"fmt"
	"github.com/SimonXming/Algorithms-practice/settings"
	"os"
	"path/filepath"
	"strings"
)

type CharSlice []string

func (c *CharSlice) Len() int {
	return len(*c)
}
func (c *CharSlice) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}
func (c *CharSlice) Less(i, j int) bool {
	// for single character, compare symbol (like < >) work too.
	return (*c)[i] < (*c)[j]
}
func (c *CharSlice) Show() {
	fmt.Printf("%v\n", *c)
}

func (c *CharSlice) ShellSort() {
	// 获取待排序的列表长度
	length := c.Len()
	h := 1

	for h < length/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		// i 是子数组开始一轮比较的起始位置
		// i 也是标记位元素，左边已排序，右边未排序
		// 从 i 开始遍历到最后，意味着在间隔为 h 的子数列编织方式下，数列变为有序状态

		for i := h; i < length; i++ {
			// i 每增大一个，意味着开始一次以 i 为起始位置，间隔为 h 的子数组排序开始了

			for j := i; j >= h && c.Less(j, j-h); j -= h {
				// 在子数组中，进行间隔为 h 的元素比较，如发现 j < j-h，则进行交换
				c.Swap(j, j-h)
			}
		}
		h = h / 3
	}
}

var (
	dataBase = settings.DataBase
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readDataFromFile(path string) (char_slice CharSlice, e error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	char_slice = make(CharSlice, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), " ")
		for _, c := range chars {
			char_slice = append(char_slice, c)
		}
	}
	return char_slice, nil
}

func newCollectionFromFile(path string) *CharSlice {
	collection, _ := readDataFromFile(path)
	return &collection
}

func main() {
	// filePath := filepath.Join(dataBase, "21elementary", "medium.txt")
	filePath := filepath.Join(dataBase, "21elementary", "tiny.txt")
	s_slict := newCollectionFromFile(filePath)
	s_slict.ShellSort()
	s_slict.Show()
}
