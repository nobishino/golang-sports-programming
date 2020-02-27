package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N, Q int
	t    *IntSegTree
	INF  int = ModPow(2, 31, 0) - 1
)

func readVariables() {
	N, Q = nextInt(), nextInt()
	t = NewIntSegTree(N, MinInt, INF)
}

func main() {
	readVariables()
	for i := 0; i < Q; i++ {
		cmd, x, y := nextInt(), nextInt(), nextInt()
		if cmd == 0 {
			t.Update(x, y)
		} else {
			fmt.Println(t.Find(x, y+1))
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

//IntSegTree は、要素型をintに限定したセグメント木です。
type IntSegTree struct {
	size        int
	unitElement int
	operation   func(int, int) int
	depth       int
	leafNum     int
	nodes       []int
}

func NewIntSegTree(size int, operation func(int, int) int, unitElement int) *IntSegTree {
	//要素を格納する木・配列の深さと大きさを計算する
	d, s := 0, 1
	for s < size {
		d++
		s *= 2
	}
	nodes := make([]int, 2*s)
	//要素を初期化する
	for i := range nodes {
		nodes[i] = unitElement
	}
	//値を返す
	return &IntSegTree{
		size:        size,
		unitElement: unitElement,
		operation:   operation,
		depth:       d + 1,
		leafNum:     s,
		nodes:       nodes,
	}
}

/*
Update は、セグメント木の要素の値を更新します。

locationで何番目の要素かを指定し(0-indexed)、valueで更新後の値を指定します。
*/
func (t *IntSegTree) Update(location, value int) {
	//leafの更新
	nodeIndex := t.leafNum - 1 + location
	t.nodes[nodeIndex] = value
	//親ノードの再計算
	for nodeIndex > 0 {
		nodeIndex = (nodeIndex - 1) / 2
		leftChild := t.nodes[2*nodeIndex+1]
		rightChild := t.nodes[2*nodeIndex+2]
		t.nodes[nodeIndex] = t.operation(leftChild, rightChild)
	}
}

//Find returns a_start * a_{start + 1} * ... * a_{end - 1}
//
//i.e. returns reduced value within [start,end)
func (t *IntSegTree) Find(start, end int) int {
	return t.helper(start, end, 0, 0, t.leafNum)
}

func (t *IntSegTree) helper(start, end, nodeIndex, left, right int) int {
	// fmt.Println("args", start, end, nodeIndex, left, right)
	//重複部分なしの場合、単位元を返す
	if right <= start || end <= left {
		return t.unitElement
	}
	//担当範囲がクエリ範囲に完全に含まれる場合、保持している値を返す
	if start <= left && right <= end {
		return t.nodes[nodeIndex]
	}
	//それ以外の場合、子要素にクエリを投げ、マージする
	mid := (left + right) / 2
	return t.operation(
		t.helper(start, end, 2*nodeIndex+1, left, mid),
		t.helper(start, end, 2*nodeIndex+2, mid, right),
	)
}
