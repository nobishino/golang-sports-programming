package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	n, k int
	h    Ints
	ans  int
)

type Ints []int

func (s Ints) Len() int {
	return len(s)
}

func (s Ints) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s Ints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	n = nextInt()
	k = nextInt()
	h = make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}
	sort.Sort(h)
	for i := k; i < n; i++ {
		ans += h[i]
	}
	fmt.Println(ans)
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
