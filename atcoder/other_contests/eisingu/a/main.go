package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	L, R, d int
)

func main() {
	defer writer.Flush()
	readVariables()
	if L%d != 0 {
		L += d - L%d
	}
	var answer int
	if R >= L {
		answer = (R-L)/d + 1
	} else {
		answer = 0
	}
	println(answer)
}

func readVariables() {
	L, R, d = nextInt(), nextInt(), nextInt()
}

/* Template */

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	Max := 1001001
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, Max), Max)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriterSize(os.Stdout, Max)
}

func println(a ...interface{}) {
	fmt.Fprintln(writer, a...)
}

func printf(format string, a ...interface{}) {
	fmt.Fprintf(writer, format, a...)
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

// MinInt returns minimum argument
func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

//MaxInt returns maximum argument
func MaxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

//AbsInt returns |x| for x
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
