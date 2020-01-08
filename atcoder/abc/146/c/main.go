package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var a, b, x int

func main() {
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&x)
	var (
		ok  int = 0
		ng  int = 1000000000
		mid int
	)
	if buyAble(ng) {
		fmt.Println(ng)
		os.Exit(0)
	}
	for math.Abs(float64(ok-ng)) > 1.0 {
		mid = (ok + ng) / 2
		if buyAble(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	fmt.Println(ok)
}

func buyAble(n int) bool {
	digit := len(strconv.Itoa(n))
	result := x >= a*n+b*digit
	return result
}
