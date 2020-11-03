package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n int

func main() {
	readVariables()
	ans := solve(n)
	fmt.Println(ans)
}

func naive(n int) string {
	var white = true
	for i := 0; i < n; i++ {
		white = !white
	}
	if white {
		return "White"
	}
	return "Black"
}

func solve(n int) string {
	if n%2 == 0 {
		return "White"
	}
	return "Black"
}

func readVariables() {
	n = nextInt()
}

/* Template */

var scanner *bufio.Scanner

func init() {
	Max := 1001001
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, Max), Max)
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
