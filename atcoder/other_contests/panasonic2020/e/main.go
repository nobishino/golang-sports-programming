package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	ss   [3]string
	conf [3][3][10000]bool //true if conflicts
)

func main() {
	readVariables()
	preCalc()
	answer := 10000
	answer = MinInt(calc(0, 1, 2), answer)
	answer = MinInt(calc(0, 2, 1), answer)
	answer = MinInt(calc(1, 0, 2), answer)
	answer = MinInt(calc(1, 2, 0), answer)
	answer = MinInt(calc(2, 0, 1), answer)
	answer = MinInt(calc(2, 1, 0), answer)
	fmt.Println(answer)
}

func calc(k, l, m int) int {
	result := 10000
	for i := 0; i < len(ss[k])+2; i++ {
		for j := i; j <= len(ss[k])+len(ss[l]); j++ {
			err := conflict(k, l, i) || conflict(k, m, j) || conflict(l, m, j-i)
			if !err {
				result = MinInt(length(k, l, m, i, j), result)
			}
		}
	}
	return result
}

func length(k, l, m, i, j int) int {
	x := len(ss[k])
	y := i + len(ss[l])
	z := j + len(ss[m])
	return MaxInt(x, MaxInt(y, z))
}

func conflict(k, l, i int) bool {
	if i < 10000 {
		return conf[k][l][i]
	} else {
		return false
	}
}

func preCalc() {
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			//k to l
			for i := 0; i < len(ss[k]); i++ {
			OUTER:
				for j := 0; i+j < len(ss[k]) && j < len(ss[l]); j++ {
					c1, c2 := ss[k][i+j], ss[l][j]
					if c1 != '?' && c2 != '?' && c1 != c2 {
						conf[k][l][i] = true
						break OUTER
					}
				}
			}
		}
	}
}

func readVariables() {
	ss[0], ss[1], ss[2] = nextStr(), nextStr(), nextStr()
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
