/*
选择排序
对于长度为 N 的数组，选择排序需要大约 N^2/2 次比较和 N 次交换。
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

func (c *CharSlice) Sort() {
	length := c.Len()
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if c.Less(j, min) {
				min = j
			}
			c.Swap(i, min)
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
	s_slict.Sort()
	s_slict.Show()
}
