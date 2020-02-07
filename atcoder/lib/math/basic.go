package math

//MinInt は、2つの整数を受け取り、最小値を返す。
func MinInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

//MaxInt は、2つの整数を受け取り、最大値を返す。
func MaxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
