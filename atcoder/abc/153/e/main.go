package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	dp   []int
	dmg  []int
	cost []int
)

const inf = 100000000000

func main() {
	H := nextInt()
	N := nextInt()
	dmg = make([]int, N)
	cost = make([]int, N)
	for i := 0; i < N; i++ {
		dmg[i] = nextInt()
		cost[i] = nextInt()
	}
	dp = make([]int, H+1)
	for i := 1; i < H+1; i++ {
		dp[i] = inf
	}
	for h := 1; h < H+1; h++ {
		for j := 0; j < N; j++ {
			if dmg[j] >= h {
				dp[h] = minInt(dp[h], cost[j])
			} else {
				dp[h] = minInt(dp[h], dp[h-dmg[j]]+cost[j])
			}
		}
	}
	fmt.Println(dp[H])
}

func minInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}
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
