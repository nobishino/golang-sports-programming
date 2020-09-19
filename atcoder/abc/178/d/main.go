package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var dp [3000]int
var s int

const mod int = 1e9 + 7

func main() {
	readVariables()
	dp[0] = 1
	for i := 1; i <= s; i++ {
		dp[i] = dp[i-1] + get(i-3)
		dp[i] %= mod
	}
	answer := get(s) - get(s-1)
	fmt.Println(answer)
}

func get(i int) int {
	if i < 0 {
		return 0
	}
	return dp[i]
}

func readVariables() {
	s = nextInt()

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
