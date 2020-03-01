package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	st   *IntSegTree
	N, Q int
	S    string
)

func sToI(s string) int {
	return int(s[0] - 'a')
}

func main() {
	readVariables()
	for i := 0; i < Q; i++ {
		q := nextInt()
		if q == 2 {
			x, y := nextInt()-1, nextInt()-1
			fmt.Println(st.Find(x, y+1))
		} else {
			l, r := nextInt()-1, nextStr()
			var m C
			m[sToI(r)] = 1
			st.Update(l, m)
		}
	}
}

func readVariables() {
	N = nextInt()
	S = nextStr()
	Q = nextInt()
	var unit C
	st = NewIntSegTree(N, add, unit)
	for i := 0; i < N; i++ {
		c := int(S[i] - 'a')
		var m C
		m[c] = 1
		st.Update(i, m)
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

type C [26]int

//IntSegTree
type IntSegTree struct {
	size        int
	unitElement C
	operation   func(C, C) C
	depth       int
	leafNum     int
	nodes       []C
}

func NewIntSegTree(size int, operation func(C, C) C, unitElement C) *IntSegTree {
	d, s := 0, 1
	for s < size {
		d++
		s *= 2
	}
	nodes := make([]C, 2*s)
	for i := range nodes {
		nodes[i] = unitElement
	}
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
Update
*/
func (t *IntSegTree) Update(location int, value C) {
	nodeIndex := t.leafNum - 1 + location
	t.nodes[nodeIndex] = value
	for nodeIndex > 0 {
		nodeIndex = (nodeIndex - 1) / 2
		leftChild := t.nodes[2*nodeIndex+1]
		rightChild := t.nodes[2*nodeIndex+2]
		t.nodes[nodeIndex] = t.operation(leftChild, rightChild)
	}
}

func add(x, y C) C {
	var result C
	for i := 0; i < 26; i++ {
		result[i] = x[i] + y[i]
	}
	return result
}

func K(c C) int {
	var result int
	for i := 0; i < 26; i++ {
		if c[i] > 0 {
			result++
		}
	}
	return result
}

//Find returns a_start * a_{start + 1} * ... * a_{end - 1}
//
//i.e. returns reduced value within [start,end)
func (t *IntSegTree) Find(start, end int) int {
	return K(t.helper(start, end, 0, 0, t.leafNum))
}

func (t *IntSegTree) helper(start, end, nodeIndex, left, right int) C {
	if right <= start || end <= left {
		return t.unitElement
	}
	if start <= left && right <= end {
		return t.nodes[nodeIndex]
	}
	mid := (left + right) / 2
	return t.operation(
		t.helper(start, end, 2*nodeIndex+1, left, mid),
		t.helper(start, end, 2*nodeIndex+2, mid, right),
	)
}
