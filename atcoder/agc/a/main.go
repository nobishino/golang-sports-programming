package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	H, W int
	b    [100][100]bool
	d    [100][100]int
)

func main() {
	readVariables()
	// fmt.Println(b[0][0], b[1][0], b[2][2])
	d[0][0] = 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if j != W-1 {
				var diff int
				if b[i][j] != b[i][j+1] {
					diff = 1
				}
				d[i][j+1] = MinInt(d[i][j]+diff, d[i][j+1])
			}
			if i != H-1 {
				var diff int
				if b[i][j] != b[i+1][j] {
					diff = 1
				}
				d[i+1][j] = MinInt(d[i][j]+diff, d[i+1][j])
			}
		}
	}
	ans := (d[H-1][W-1] + 1) / 2
	// for i := 0; i < H; i++ {
	// 	fmt.Println(d[i][:W])
	// }
	if !b[0][0] && !b[H-1][W-1] {
		ans++
	}
	fmt.Println(ans)
}

func readVariables() {
	H, W = nextInt(), nextInt()
	for i := 0; i < H; i++ {
		r := nextStr()
		for j := 0; j < W; j++ {
			if r[j] == '.' {
				b[i][j] = true
			}
			d[i][j] = 1234567890
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
