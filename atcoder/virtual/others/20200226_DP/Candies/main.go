package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	dp           [101][100001]int
	N, K, answer int
	a            [100]int
	MOD          = 1000000007
)

func readVariables() {
	N, K = nextInt(), nextInt()
	for i := 0; i < N; i++ {
		a[i] = nextInt()
	}
}

func main() {
	readVariables()
	dp[0][0] = 1
	for i := 0; i < N; i++ {
		for j := 0; j < K+1; j++ {
			if j-a[i]-1 >= 0 {
				dp[i+1][j] = dp[i][j] - dp[i][j-a[i]-1]
			} else {
				dp[i+1][j] = dp[i][j]
			}
			dp[i+1][j] = Mod(dp[i+1][j], MOD)
		}
	}
	for i := 0; i < N; i++ {
		for j := 1; j < K+1; j++ {
			dp[N][j] += dp[N][j-1]
			dp[N][j] = Mod(dp[N][j], MOD)
		}
	}
	answer = dp[N][K]
	// for i := 0; i < N+1; i++ {
	// 	for j := 0; j < K+1; j++ {
	// 		fmt.Print(dp[i][j], ",")
	// 	}
	// 	fmt.Println()
	// }
	fmt.Println(answer)
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

func Mod(x, y int) int {
	if x >= 0 {
		return x % y
	} else {
		return y + x%y
	}
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
