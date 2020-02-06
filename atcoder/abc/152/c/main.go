package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N int
	P []int
)

func readVariables() {
	N = nextInt()
	P = make([]int, N)
	for i := 0; i < N; i++ {
		P[i] = nextInt()
	}
}

func main() {
	readVariables()
	minValue := N + 100 //INF
	answer := 0
	for _, v := range P {
		if v <= minValue {
			answer++
			minValue = v
		}
	}
	fmt.Println(answer)
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
