package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	H, W, K    int
	black, red [][]bool
)

func main() {
	defer writer.Flush()
	readVariables()
	var answer int
	for br := 0; br < 1<<H; br++ {
		for bc := 0; bc < 1<<W; bc++ {
			clearRed()
			for i := 0; i < H; i++ {
				if 1<<i&br > 0 {
					setRowRed(i)
				}
			}
			for j := 0; j < W; j++ {
				if 1<<j&bc > 0 {
					setColRed(j)
				}
			}
			k := count()
			if k == K {
				answer++
			}
		}
	}
	println(answer)
}

func count() int {
	var result int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if black[i][j] && !red[i][j] {
				result++
			}
		}
	}
	return result
}

func setRowRed(i int) {
	for j := 0; j < W; j++ {
		red[i][j] = true
	}
}
func setColRed(j int) {
	for i := 0; i < H; i++ {
		red[i][j] = true
	}
}

func clearRed() {
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			red[i][j] = false
		}
	}
}

func readVariables() {
	H, W, K = nextInt(), nextInt(), nextInt()
	black = make([][]bool, H)
	red = make([][]bool, H)
	for i := 0; i < H; i++ {
		black[i] = make([]bool, W)
		red[i] = make([]bool, W)
	}
	for i := 0; i < H; i++ {
		row := nextStr()
		for j := 0; j < W; j++ {
			black[i][j] = row[j] == '#'
		}
	}
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
