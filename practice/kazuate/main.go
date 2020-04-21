package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var answers []int
	for i := 0; i < 1000000; i++ {
		if satisfy(i) {
			answers = append(answers, i)
		}
	}
	fmt.Println(len(answers)) //3
}

func digits(x int) int {
	return len(strconv.Itoa(x))
}

func satisfy(x int) bool {
	return (digitsSum(x) == 22) && (digits(x) == 5) && (countPhi(x) == 15)
}

func digitsSum(x int) int {
	var ans int
	for x > 0 {
		ans += x % 10
		x /= 10
	}
	return ans
}

func countPhi(x int) int {
	bound := int(math.Sqrt(float64(x))) + 2
	divisible := make(map[int]bool)
	for i := 1; i < bound; i++ {
		if x%i == 0 {
			divisible[i] = true
			divisible[x/i] = true
		}
	}
	return len(divisible)
}
