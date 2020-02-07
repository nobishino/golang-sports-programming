package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N, answer int
	cnt       [10][10]int
)

func readVariables() {
	N = nextInt()
}

func main() {
	readVariables()
	preCalc()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			answer += cnt[i][j] * cnt[j][i]
		}
	}
	fmt.Println(answer)
}

func preCalc() {
	for i := 1; i <= N; i++ {
		head, tail := getHeadTail(i)
		cnt[head][tail]++
	}
}

func getHeadTail(x int) (head, tail int) {
	strX := strconv.Itoa(x)
	head, _ = strconv.Atoi(string(strX[0]))
	tail, _ = strconv.Atoi(string(strX[len(strX)-1]))
	return
}

func count(head, tail, maxValue int) (result int) {
	if head == 0 {
		return
	}
	if head == tail && head <= maxValue {
		result++
	}
	if head*10+tail <= maxValue {
		result++
	}
	if head*100+tail > maxValue {
		return
	}
	ok := 0
	ng := maxValue
	// fmt.Println(ok, ng)
	for AbsInt(ok-ng) > 1 {
		var count int
		count++
		mid := (ok + ng) / 2
		value := tail + mid*10 + head*powInt(10, digits(mid)+1)
		if count > 10 {
			fmt.Println(ok, ng, mid, maxValue, value)
		}
		if value <= maxValue {
			ok = mid
		} else {
			ng = mid
		}
	}
	result += ok + 1
	return
}

func digits(x int) int {
	return len(strconv.Itoa(x))
}

func powInt(x, p int) (result int) {
	result = 1
	for i := 0; i < p; i++ {
		result *= x
	}
	return
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
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
