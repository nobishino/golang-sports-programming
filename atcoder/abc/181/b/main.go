package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readVariables()
	ans := solve(input)
	fmt.Println(ans)
}

type in struct {
	n int
	a []int
	b []int
}

func readVariables() in {
	n := nextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		b[i] = nextInt()
	}
	return in{
		n: n,
		a: a,
		b: b,
	}
}

func solve(input in) int {
	n, a, b := input.n, input.a, input.b
	var ans int
	for i := 0; i < n; i++ {
		diff := ((a[i] + b[i]) * (b[i] - a[i] + 1)) / 2
		ans += diff
	}
	return ans
}

func naive(input in) int {
	n, a, b := input.n, input.a, input.b
	var ans int
	for i := 0; i < n; i++ {
		for v := a[i]; v <= b[i]; v++ {
			ans += v
		}
	}
	return ans
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
