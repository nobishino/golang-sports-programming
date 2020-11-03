package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	n    int
	x, y []int
)

func main() {
	readVariables()
	fmt.Println(solve(n, x, y))
}

func solve(n int, x, y []int) string {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a1, a2 := diff(x[i], y[i], x[j], y[j])
			for k := j + 1; k < n; k++ {
				b1, b2 := diff(x[i], y[i], x[k], y[k])
				d := determinant(a1, a2, b1, b2)
				if d == 0 {
					return "Yes"
				}
			}
		}
	}
	return "No"
}

func determinant(x1, y1, x2, y2 int) int {
	return x1*y2 - x2*y1
}

func diff(x1, y1, x2, y2 int) (int, int) {
	return x2 - x1, y2 - y1
}

func readVariables() {
	n = nextInt()
	x = make([]int, n)
	y = make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
		y[i] = nextInt()
	}
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
