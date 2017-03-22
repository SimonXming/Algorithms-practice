package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/SimonXming/Algorithms-practice/lib"
	"github.com/SimonXming/Algorithms-practice/settings"
	"os"
	"path/filepath"
	"sort"
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
		b_index := sort.SearchInts(numbers, -numbers[a])
		if b_index != numCount {
			if sum := numbers[a] + numbers[b_index]; sum == 0 && a < b_index {
				count++
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
	count := twoSumCount(lines)
	end := time.Now()
	delta := end.Sub(start)

	fmt.Printf("count: %d, time: %s \n", count, delta)
}
