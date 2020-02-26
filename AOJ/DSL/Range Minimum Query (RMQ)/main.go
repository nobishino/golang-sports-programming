package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	N, Q     int
	exponent int
	nodes    [262144]int
	INF      int = ModPow(2, 31, 0) - 1
)

func readVariables() {
	N, Q = nextInt(), nextInt()
	InitSegTree(N)
}
func InitSegTree(size int) {
	exponent = int(math.Ceil(math.Log2(float64(size))))
	// fmt.Println("exponent = ", exponent)
	bound := ModPow(2, exponent, 0) * 2
	for i := range nodes {
		if i > bound {
			break
		}
		nodes[i] = INF
	}
}

func update(loc, value int) {
	index := ModPow(2, exponent, 0) - 1 + loc
	nodes[index] = value
	for index > 0 {
		index = (index - 1) / 2
		nodes[index] = MinInt(nodes[2*index+1], nodes[2*index+2])
	}
}
func find(start, term int) int {
	return helpFind(start, term, 0, 0, ModPow(2, exponent, 0))
}
func helpFind(start, term, nodeIndex, nodeStart, nodeEnd int) int {
	if term < nodeStart || nodeEnd <= start {
		return INF
	}
	if term < nodeEnd-1 || start > nodeStart {
		mid := nodeStart + (nodeEnd-nodeStart)/2
		return MinInt(
			helpFind(start, term, 2*nodeIndex+1, nodeStart, mid),
			helpFind(start, term, 2*nodeIndex+2, mid, nodeEnd),
		)
	}
	// fmt.Println(nodeIndex)
	return nodes[nodeIndex]
}

func main() {
	readVariables()
	for i := 0; i < Q; i++ {
		cmd, x, y := nextInt(), nextInt(), nextInt()
		if cmd == 0 {
			update(x, y)
		} else {
			fmt.Println(find(x, y))
		}
	}
}

/* 以下、テンプレート*/

var scanner *bufio.Scanner

func init() {
	Max := 10010010
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

// MinInt は、2つの整数を受け取り、最小値を返します。
func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

//MaxInt は、2つの整数を受け取り、最大値を返します。
func MaxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

//AbsInt は、整数の絶対値を返します。
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//ModPow は整数の累乗関数で、剰余もサポートします。
//base^exponentの値をmoduleで割った余りを返します。
//moduloが1以下の場合には、剰余演算をしません。
//baseが負の値である場合には、返す値が負になることがあります。
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

//Gcd は、引数の整数全ての最大公約数を返します。
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

//Lcm は、与えられた整数の最小公倍数を返します。
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
