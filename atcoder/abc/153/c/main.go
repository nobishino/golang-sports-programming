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
	h    []int
	ans  int
)

func main() {
	n = nextInt()
	k = nextInt()
	h = make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}
	sort.Ints(h)
	for i := 0; i < n-k; i++ {
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
