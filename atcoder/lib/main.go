//このファイルの内容をAtCoder用テンプレートとして用いる。
package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {

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

//MinInt は、2つの整数を受け取り、最小値を返す。
func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

//MaxInt は、2つの整数を受け取り、最大値を返す。
func MaxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
