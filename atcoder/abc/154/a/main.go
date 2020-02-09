package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	S, T, U string
	A, B    int
)

func readVariables() {
	S, T = nextStr(), nextStr()
	A, B = nextInt(), nextInt()
	U = nextStr()
}

func main() {
	readVariables()
	if U == S {
		A--
	} else {
		B--
	}
	fmt.Println(A, B)
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
