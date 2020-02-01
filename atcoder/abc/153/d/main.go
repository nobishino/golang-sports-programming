package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	H := nextInt()
	answer := solve(H)
	fmt.Println(answer)
}

func solve(h int) int {
	if h == 0 {
		return 0
	}
	if h == 1 {
		return 1
	}
	return 1 + 2*solve(h/2)
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
