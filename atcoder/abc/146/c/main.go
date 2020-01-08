package main

import (
	"fmt"
	"math"
	"strconv"
)

var a, b, x int

func main() {
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&x)
	var (
		ok     int = 0
		ng     int = 1000000000
		answer int
		mid    int
	)
	if buyAble(ng) {
		answer = ng
	} else {
		for math.Abs(float64(ok-ng)) > 1.0 {
			mid = (ok + ng) / 2
			if buyAble(mid) {
				ok = mid
			} else {
				ng = mid
			}
		}
		answer = ok
	}
	fmt.Println(answer)
}

func buyAble(n int) bool {
	digit := len(strconv.Itoa(n))
	result := x >= a*n+b*digit
	return result
}
