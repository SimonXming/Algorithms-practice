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

var readIDListCount int = 0

type Pair struct {
	Items []int
}

type UF struct {
	count int
	id    []int
	sz    []int
}

func newUF(N int) *UF {
	ids := make([]int, 0)
	size := make([]int, 0)
	for i := 0; i < N; i++ {
		ids = append(ids, i)
		size = append(size, 1)
	}
	return &UF{count: N, id: ids, sz: size}
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) Find(p int) int {
	// 寻找 p 所处的分量
	// 沿着树结构寻找根节点
	// 根节点的触点值代表了这个分量
	for p != uf.id[p] {
		readIDListCount++
		p = uf.id[p]
	}
	return p
}

func (uf *UF) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) QuickUnion(p int, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)

	if pRoot == qRoot {
		// This mean q and p in the same 分量
		return
	}
	// 根据分量的大小判断根节点的指向 q -> p 或者 p -> q
	// 这就是说 q p 所处的两个分量联合在一起
	if uf.sz[pRoot] < uf.sz[qRoot] {
		readIDListCount++
		uf.id[pRoot] = uf.id[qRoot]
		uf.sz[qRoot] += uf.id[pRoot]
	} else {
		readIDListCount++
		uf.id[qRoot] = uf.id[pRoot]
		uf.sz[pRoot] += uf.id[qRoot]
	}
	uf.count--
}

func newPair(s string) Pair {
	nums := strings.Split(s, " ")
	items := make([]int, 0)
	for _, num := range nums {
		item, _ := strconv.Atoi(num)
		items = append(items, item)
	}
	return Pair{items}
}

var (
	dataBase     = settings.DataBase
	filenameFlag = flag.String("filename", "", "Description")
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readDataFromFile(path string) (N int, pairs []Pair, e error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := newPair(scanner.Text())
		if len(t.Items) == 1 {
			N = t.Items[0]
		} else {
			pairs = append(pairs, t)
		}
	}
	return N, pairs, nil
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

	n, pairs, err := readDataFromFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	start := time.Now()

	uf := newUF(n)

	for _, pair := range pairs {
		p, q := pair.Items[0], pair.Items[1]
		// fmt.Printf("Check connected, %d %d \r\n%t \r\n", p, q, uf.Find(p) == uf.Find(q))
		if uf.Connected(p, q) {
			continue
		}
		// fmt.Println("Before UNION", uf.id)
		uf.QuickUnion(p, q)
		// fmt.Println(p, " + ", q)
		// fmt.Println("After UNION", uf.id)
	}
	end := time.Now()
	delta := end.Sub(start)

	fmt.Printf("%d components, time: %s \n", uf.Count(), delta)
	fmt.Println("readIDListCount: %d", readIDListCount)
}
