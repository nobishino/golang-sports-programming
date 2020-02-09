package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N int
	C map[int]int = make(map[int]int)
)

func readVariables() {
	N = nextInt()
}

func main() {
	readVariables()
	var answer bool
	answer = true
	for i := 0; i < N; i++ {
		v := nextInt()
		if C[v] > 0 {
			answer = false
			break
		}
		C[v]++
	}
	// fmt.Println(C)
	if answer {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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
