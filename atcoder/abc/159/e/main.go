package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	H, W, K int
	b       [][]int
	cSum    [][]int
)

func main() {
	readVariables()
	// fmt.Println(cSum)
	answer := H + W
WALL:
	for bit := 0; bit < 1<<uint(H-1); bit++ {
		cand := 0
		// fmt.Println(bit, count(bit))
		left := 0
		right := 0
		for right < W {
			c := check(left, right, bit)
			if c {
				right++
				continue
			} else if left == right {
				continue WALL
			} else {
				cand++
				left = right
			}
		}
		cand += count(bit)
		answer = MinInt(cand, answer)
	}
	fmt.Println(answer)
}

//[l,r]が与えられたwbitの元で条件を満たすかどうか
func check(l, r, wbit int) bool {
	s := 0
	for i := 0; i < H; i++ {
		if walled(i-1, wbit) {
			s = 0
		}
		s += getSum(l, r, i, i)
		if s > K {
			return false
		}
	}
	return true
}

//立っているビット数
func count(wbit int) int {
	result := 0
	for wbit > 0 {
		result += wbit % 2
		wbit /= 2
	}
	return result
}

//i列目の「後」がwalledである時true
func walled(i, wbit int) bool {
	if i < 0 {
		return false
	}
	return wbit&(1<<uint(i)) > 0
}

func getSum(left, right, top, bottom int) int {
	return cSum[bottom+1][right+1] - cSum[bottom+1][left] - cSum[top][right+1] + cSum[top][left]
}

func readVariables() {
	H, W, K = nextInt(), nextInt(), nextInt()
	b = make([][]int, H)
	cSum = make([][]int, H+1)
	for i := 0; i < H; i++ {
		b[i] = make([]int, W)
	}
	for i := 0; i < H+1; i++ {
		cSum[i] = make([]int, W+1)
	}
	for i := 0; i < H; i++ {
		row := nextStr()
		for j := 0; j < W; j++ {
			if row[j] == '1' {
				b[i][j] = 1
			}
			cSum[i+1][j+1] = cSum[i][j+1] + cSum[i+1][j] + b[i][j] - cSum[i][j]
		}
	}
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

func nextStr() string {
	if !scanner.Scan() {
		panic("No more token.")
	}
	return scanner.Text()
}

// MinInt returns minimum argument
func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

//MaxInt returns maximum argument
func MaxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

//AbsInt returns |x| for x
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//ModPow calculates integer power with modulo operation
//if modulo <= 1, it powers w/o module operation
//if base < 0, return value might be negative too.
func ModPow(base, exponent, modulo int) (result int) {
	result = 1
	for exponent > 0 {
		if exponent%2 == 1 {
			result *= base
			if modulo > 1 {
				result %= modulo
			}
		}
		base *= base
		if modulo > 1 {
			base %= modulo
		}
		exponent /= 2
	}
	return
}

//Gcd
func Gcd(vals ...int) (result int) {
	if len(vals) == 0 {
		return
	}
	result = vals[0]
	for i := 1; i < len(vals); i++ {
		result = gcd(result, vals[i])
	}
	return
}

func gcd(x, y int) int {
	x, y = AbsInt(x), AbsInt(y)
	for y > 0 {
		x, y = y, x%y
	}
	return x
}

//Lcm
func Lcm(vals ...int) (result int) {
	if len(vals) == 0 {
		return
	}
	result = vals[0]
	for i := 1; i < len(vals); i++ {
		result = lcm(result, vals[i])
	}
	return
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
