package main

import (
	"bufio"
	"fmt"
	"github.com/SimonXming/Algorithms-practice/settings"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var dataBase = settings.DataBase

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

func twoSumCount(numbers []int) int {
	count := 0
	numCount := len(numbers)
	for a := 0; a < numCount; a++ {
		for b := a + 1; b < numCount; b++ {
			if sum := numbers[a] + numbers[b]; sum == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	filename := "32Kints.txt"
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
