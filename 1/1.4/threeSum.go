package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const dataBase = "/Users/simon/Code/go/src/github.com/SimonXming/Algorithms-practice/data"

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
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

func threeSumCount(numbers []int) int {
	count := 0
	numCount := len(numbers)
	for a := 0; a < numCount; a++ {
		for b := a + 1; b < numCount; b++ {
			for c := b + 1; c < numCount; c++ {
				if sum := numbers[a] + numbers[b] + numbers[c]; sum == 0 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	filename := "2Kints.txt"
	filePath := filepath.Join(dataBase, filename)

	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	start := time.Now()
	count := threeSumCount(lines)
	end := time.Now()
	delta := end.Sub(start)

	fmt.Printf("count: %d, time: %s \n", count, delta)
}