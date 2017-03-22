package main

import (
	"bufio"
	"fmt"
	"github.com/SimonXming/Algorithms-practice/lib"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const dataBase = "/Users/simon/Code/go/src/github.com/SimonXming/Algorithms-practice/data"

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (lib.IntSlice, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines lib.IntSlice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		i, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

func twoSumCount(numbers lib.IntSlice) int {
	count := 0
	sort.Sort(numbers)
	numCount := len(numbers)
	for a := 0; a < numCount; a++ {
		// fmt.Println(a)
		b_index := sort.SearchInts(numbers, -numbers[a])
		// fmt.Println(b_index)

		if sum := numbers[a] + numbers[b_index]; sum == 0 && a < b_index {
			count++
		}
	}
	return count
}

func main() {
	filename := "1Kints.txt"
	filePath := filepath.Join(dataBase, filename)

	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	start := time.Now()
	count := twoSumCount(lines)
	end := time.Now()
	delta := end.Sub(start)

	fmt.Printf("count: %d, time: %s \n", count, delta)
}
