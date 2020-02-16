package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N, W, K, V int
	c, v       [1000]int
	answer     [1000]int
)

func readVariables() {
	N = nextInt()
	W = nextInt()
	K = nextInt()
	V = nextInt()
	for i := 0; i < N; i++ {
		c[i], v[i] = nextInt(), nextInt()
	}
}

func main() {
	readVariables()
	solve()
	p()
}
func p() {
	for _, v := range answer {
		fmt.Println(v)
	}
}

func solve() {
	for i := 0; i < N; i++ {
		answer[i] = (i + 1) % W
	}
}

/* 以下、テンプレート*/

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 10000000), 10000000)
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
