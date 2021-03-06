package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	k, q    int
	d, md   [5000]int
	n, x, m int
)

func readVariables() {
	k, q = nextInt(), nextInt()
	for i := 0; i < k; i++ {
		d[i] = nextInt()
	}
}

func readQParam() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
}

func query(n, x, m int) int {
	// fmt.Println(n, x, m)
	x %= m
	calcMod(m)
	sum := calcPeriodSum(k)
	q := (n - 1) / k
	r := (n - 1) % k
	last := x + q*sum
	for i := 0; i < r; i++ {
		last += md[i]
	}
	// fmt.Println(sum)
	// fmt.Println(md[0:k])
	return n - 1 - last/m
}

func calcMod(m int) {
	for i := 0; i < k; i++ {
		md[i] = d[i] % m
		if md[i] == 0 {
			md[i] = m
		}
	}
}

func calcPeriodSum(k int) int {
	result := 0
	for i := 0; i < k; i++ {
		result += md[i]
	}
	return result
}

func main() {
	readVariables()
	for i := 0; i < q; i++ {
		n, x, m = readQParam()
		fmt.Println(query(n, x, m))
	}
}

/* 以下、テンプレート*/

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
