package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N, Q  int
	links [154321][]Link
	MOD   int = 1000000007
)

type Link struct {
	dest int
	m    int
}

func readVariables() {
	N = nextInt()
	for i := 0; i < N-1; i++ {
		u, v, c := nextInt(), nextInt(), nextInt()
		links[u] = append(links[u], Link{v, c})
		links[v] = append(links[v], Link{u, c})
	}
	Q = nextInt()
}

func main() {
	readVariables()
	for q := 0; q < Q; q++ {
		m, p, x := nextInt(), nextInt(), nextInt()
		a := query(m, p, x)
		fmt.Println(a)
	}
}

func query(m, p, x int) int {
	vals := make([]int, N+1)
	for i := range vals {
		vals[i] = -1
	}
	vals[m] = 1
	var q Queue = make([]int, 0)
	q.Offer(m)
	for len(q) > 0 {
		v := q.Pop()
		for _, link := range links[v] {
			w := link.dest
			c := link.m
			if vals[w] != -1 {
				continue
			}
			vals[w] = c * vals[v]
			vals[w] %= MOD
			q.Offer(w)
		}
	}
	// fmt.Println(vals)
	return (x * vals[p]) % MOD
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

type Queue []int

func (q *Queue) Offer(x int) {
	*q = append(*q, x)
}
func (q *Queue) Pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
