package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N, Q int
	A    []int
	S    []int
)

func main() {
	readVariables()
	for i := 1; i < N; i++ {
		A[i] = Gcd(A[i-1], A[i])
	}
	for q := 0; q < Q; q++ {
		if x := Gcd(S[q], A[N-1]); x != 1 {
			fmt.Println(x)
			continue
		}
		ok, ng := N-1, -1
		for AbsInt(ok-ng) > 1 {
			mid := (ok + ng) / 2
			if Gcd(S[q], A[mid]) == 1 {
				ok = mid
			} else {
				ng = mid
			}
		}
		fmt.Println(ok + 1)
	}
}

func readVariables() {
	N, Q = nextInt(), nextInt()
	A = make([]int, N)
	S = make([]int, Q)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}
	for i := 0; i < Q; i++ {
		S[i] = nextInt()
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
