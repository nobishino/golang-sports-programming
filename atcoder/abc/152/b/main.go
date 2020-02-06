package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var a, b int

func readVariables() {
	a, b = nextInt(), nextInt()
}

func main() {
	readVariables()
	if a > b {
		a, b = b, a
	}
	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
	fmt.Println()
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
