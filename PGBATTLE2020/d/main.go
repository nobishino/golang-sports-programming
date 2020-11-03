package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	n, m int
	link [][]int
	inv  []int
)

func main() {
	readVariables()
	inv[0] = 0
	var q Q = Q{0}
	for !q.empty() {
		v := q.pop()
		// fmt.Println(q, v, link[v])
		for _, w := range link[v] {
			if inv[w] != -1 {
				continue
			}
			// fmt.Println("reach", w, "from", v)
			inv[w] = v
			q.append(w)
		}
	}
	fmt.Println(ans(inv))
}

func ans(inv []int) string {
	if inv[n-1] == -1 {
		return "-1"
	}
	var a []int
	v := n - 1
	for {
		a = append(a, v)
		if v == 0 {
			break
		}
		v = inv[v]
	}
	var r []string
	for i := len(a) - 1; i >= 0; i-- {
		r = append(r, strconv.Itoa(a[i]+1))
	}
	return strings.Join(r, " ")
}

type Q []int

func (q *Q) append(v int) {
	*q = append(*q, v)
}

func (q *Q) pop() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Q) empty() bool {
	return len(*q) == 0
}

func readVariables() {
	n = nextInt()
	m = nextInt()
	link = make([][]int, n)
	for i := 0; i < n; i++ {
		link[i] = make([]int, 0)
	}
	inv = make([]int, n)
	for i := 0; i < m; i++ {
		// make edge 0-indexed
		v := nextInt() - 1
		w := nextInt() - 1
		link[v] = append(link[v], w)
	}
	for i := range link {
		sort.Slice(link[i], func(x, y int) bool {
			return link[i][x] < link[i][y]
		})
	}
	for i := 0; i < n; i++ {
		inv[i] = -1
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
