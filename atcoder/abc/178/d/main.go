package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod int = 1e9 + 7

var s int

func main() {
	readVariables()
	fmt.Println(solve2(s))
	for i := 1; i <= s; i++ {
		if solve1(s) != solve2(s) {
			fmt.Println(s, solve1(s), solve2(s))
		}
	}
}

// O(S^2)
func solve1(s int) int {
	var dp [3000]int
	dp[0] = 1
	for i := 1; i <= s; i++ {
		for j := 3; i-j >= 0; j++ {
			dp[i] += dp[i-j]
			dp[i] %= mod
		}
	}
	answer := dp[s]
	return answer
}

// intended O(S)
func solve2(s int) int {
	var dp [3000]int
	dp[0] = 1
	get := func(i int) int {
		if i < 0 {
			return 0
		}
		return dp[i]
	}
	for i := 1; i <= s; i++ {
		dp[i] = dp[i-1] + get(i-3)
		dp[i] %= mod
	}
	answer := get(s) - get(s-1)
	return answer
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
