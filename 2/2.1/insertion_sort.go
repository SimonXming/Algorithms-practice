/*
插入排序
对于随机排列的长度 n 且主键不重复的数组：
平均情况下插入排序需要 n^2/4 次比较和 n^2/4 次交换；
最坏情况 n^2/2 次比较和 n^2/2 次交换；
最好情况 n-1 次比较和 0 次交换。

插入排序对部分有序的数组很有效：
* 数组中每个元素距离它的最终位置都不远
* 一个有序的大数组接一个小数组
* 数组中只有几个元素的位置不正确

事实上，当倒置的元素很少时，插入排序很可能比本章中其他任何算法都快。

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
func (c *CharSlice) RigheSwap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
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

func (c *CharSlice) InsertionSort() {
	// 获取待排序的列表长度
	length := c.Len()
	// 逐次遍历列表
	for i := 0; i < length; i++ {
		// 设置标记位元素, 标记位元素左侧是已经排序的，标记位元素右侧是未排序的
		// 每次标记位右移，就是引入一个新的未排序元素 j，并在已经排序的元素中找一个合适的位置插入
		for j := i; j > 0 && c.Less(j, j-1); j-- {
			// 满足上述时，j 元素和前面的元素交换一下位置
			// 直到不满足条件，j 元素此时已处于已排序的位置。
			c.Swap(j, j-1)
			// 要大幅提高插入排序的速度，只需要在内循环中将较大的元素向右移动而不是交换位置
			// 这样访问数组的次数就能减半
		}
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
	filePath := filepath.Join(dataBase, "21elementary", "tiny.txt")
	// filePath := filepath.Join(dataBase, "21elementary", "tiny.txt")
	s_slict := newCollectionFromFile(filePath)
	s_slict.InsertionSort()
	s_slict.Show()
}
