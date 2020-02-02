package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	N, D, A int
	ms      PairList
	answer  int
)

func main() {
	N = nextInt()
	D = nextInt()
	A = nextInt()
	ms = make([]Pair, N)
	for i := 0; i < N; i++ {
		x := nextInt()
		h := nextInt()
		ms[i] = Pair{x, h}
	}
	sort.Sort(ms)
	bit := NewBIT(N)
	for i := 0; i < N; i++ {
		dmg := bit.Accumulate(i)
		h := ms[i].Health
		h -= dmg
		if h <= 0 {
			continue
		}
		attack := (h + A - 1) / A
		answer += attack
		newDmg := attack * A
		bit.Add(i, newDmg)
		upper := minIndexLargerThan(ms[i].Loc + 2*D)
		bit.Add(upper, -newDmg)
	}
	fmt.Println(answer)
}

func minIndexLargerThan(loc int) (result int) {
	ok := N
	ng := -1
	for math.Abs(float64(ok-ng)) > 1.0 {
		mid := (ok + ng) / 2
		if ms[mid].Loc > loc {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

type PairList []Pair

func (pl PairList) Len() int {
	return len(pl)
}

func (pl PairList) Less(i, j int) bool {
	return pl[i].Loc < pl[j].Loc
}

func (pl PairList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

type Pair struct {
	Loc    int
	Health int
}

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}
func nextInt() int {
	if !scanner.Scan() {
		panic("No more token.")
	}
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("nextInt(): cannot convert to int: " + scanner.Text())
	}
	return num
}

type BIT struct {
	Nodes []int
	Size  int
}

func NewBIT(size int) *BIT {
	bit := new(BIT)
	bit.Size = size
	bit.Nodes = make([]int, size+1)
	return bit
}

func (bit *BIT) Add(location, value int) {
	for i := location + 1; i <= bit.Size; i += (i & -i) {
		bit.Nodes[i] += value
	}
}

func (bit *BIT) Sum(start, end int) int {
	return bit.Accumulate(end) - bit.Accumulate(start-1)
}

func (bit *BIT) Accumulate(end int) (result int) {
	if end >= bit.Size {
		end = bit.Size - 1
	}
	for i := end + 1; i > 0; i -= i & -i {
		result += bit.Nodes[i]
	}
	return
}
