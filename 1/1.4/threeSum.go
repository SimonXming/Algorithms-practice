package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/SimonXming/Algorithms-practice/settings"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	dataBase     = settings.DataBase
	filenameFlag = flag.String("filename", "", "Description")
)

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

func init() {
	flag.StringVar(filenameFlag, "f", "", `filename`)
}

func main() {
	flag.Parse()
	if *filenameFlag == "" {
		fmt.Println("No data filename provide.")
		os.Exit(2)
	}

	filePath := filepath.Join(dataBase, *filenameFlag)

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
