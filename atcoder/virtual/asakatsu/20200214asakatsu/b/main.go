package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Sl []string

var (
	N, L int
	s    Sl
)

func (sl Sl) Len() int {
	return len(sl)
}

func (sl Sl) Less(i, j int) bool {
	return sl[i] < sl[j]
}

func (sl Sl) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

func readVariables() {
	N, L = nextInt(), nextInt()
	s = make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = nextStr()
	}
}

func main() {
	readVariables()
	sort.Sort(s)
	// fmt.Println(s)
	for i := 0; i < N; i++ {
		fmt.Print(s[i])
	}
	fmt.Println()
}

/* 以下、テンプレート*/

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

//nextInt converts next token from stdin and returns integer value.
//nextInt panics when conversion into an integer fails.
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

func nextStr() string {
	if !scanner.Scan() {
		panic("No more token.")
	}
	return scanner.Text()
}
