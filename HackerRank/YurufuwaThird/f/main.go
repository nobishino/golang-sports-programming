package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N              int
	A, B, C, D     [254321]int
	Sa, Sb, Sc, Sd [254321]int
	L, R           [254321]int
)

func readVariables() {
	N = nextInt()
	for i := 1; i <= N; i++ {
		A[i] = nextInt()
		Sa[i] = Sa[i-1] + A[i]
	}
	for i := 1; i <= N; i++ {
		B[i] = nextInt()
		Sb[i] = Sb[i-1] + B[i]
	}
	for i := 1; i <= N; i++ {
		C[i] = nextInt()
		Sc[i] = Sc[i-1] + C[i]
	}
	for i := 1; i <= N; i++ {
		D[i] = nextInt()
		Sd[i] = Sd[i-1] + D[i]
	}
}

func main() {
	readVariables()
	left()
	// fmt.Println(A[1 : N+1])
	// fmt.Println(B[1 : N+1])
	// fmt.Println(L[0 : N+1])
	right()
	answer := 0
	for j := 2; j <= N-2; j++ {
		a := L[j] + R[j]
		answer = MaxInt(a, answer)
	}
	fmt.Println(answer)
}

func left() {
	L[2] = A[1] + B[2]
	for i := 2; i < N; i++ {
		L[i+1] = MaxInt(L[i], Sa[i]) + B[i+1]
	}
}
func right() {
	R[N-2] = C[N-1] + D[N]
	for i := N - 2; i > 0; i-- {
		R[i-1] = MaxInt(R[i], Sd[N]-Sd[i]) + C[i]
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
