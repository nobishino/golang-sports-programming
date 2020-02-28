package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	N, M, X int
	T       [14321]int
	links   [14321][]Link
	dp      [14321][3][201]int
	INF     int = 123456789012345
	pq      PQ
)

func main() {
	readVariables()
	dijkstra()
	fmt.Println(answer())
	// fmt.Println(dp[N][0], dp[N][2])
}
func answer() int {
	result := INF
	for i := 0; i < X+1; i++ {
		result = MinInt(dp[N][0][i], result)
		result = MinInt(dp[N][2][i], result)
	}
	return result
}
func dijkstra() {
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Point)
		// vt := T[v.loc]
		t := v.t
		x := v.x
		for _, l := range links[v.loc] {
			w := l.dest
			wt := T[w]
			if wt != 1 && t != wt && x-l.weight > 0 {
				continue //遷移不可能
			}
			var newx, newt int
			if wt == 1 {
				newx = MaxInt(x-l.weight, 0)
				newt = t
			}
			if wt != 1 {
				newx = X
				newt = wt
			}
			newValue := v.value + l.weight
			if newValue < dp[w][newt][newx] {
				dp[w][newt][newx] = v.value + l.weight
				next := Point{w, v.value + l.weight, newt, newx}
				heap.Push(&pq, &next)
			}
		}
	}
}
func readVariables() {
	N, M, X = nextInt(), nextInt(), nextInt()
	for i := 1; i < N+1; i++ {
		v := nextInt()
		T[i] = v
	}
	for i := 0; i < M; i++ {
		a, b, d := nextInt(), nextInt(), nextInt()
		links[a] = append(links[a], Link{b, d})
		links[b] = append(links[b], Link{a, d})
	}
	for i := 0; i < N+1; i++ {
		for t := 0; t < 3; t++ {
			for x := 0; x <= X; x++ {
				dp[i][t][x] = INF
			}
		}
	}
	dp[1][0][X] = 0
	info := Point{1, 0, 0, X}
	heap.Init(&pq)
	heap.Push(&pq, &info)
}

/*
Priority Queue Sample Code
Usage
var pq PQ
heap.Init(&pq)
heap.Push(&pq, &point)
q := heap.Pop(&pq).(*Point)
*/
type Point struct {
	loc   int
	value int
	t     int
	x     int
}

type PQ []*Point

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].value < pq[j].value }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) {
	item := x.(*Point)
	*pq = append(*pq, item)
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

type Link struct {
	dest   int
	weight int
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

type Queue []int

func (q *Queue) Size() int {
	return len(*q)
}
func (q *Queue) Offer(v int) {
	*q = append(*q, v)
}
func (q *Queue) Pop() int {
	if q.Size() == 0 {
		panic("Pop called on empty Queue")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}
