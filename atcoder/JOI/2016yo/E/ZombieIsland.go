package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var (
	N, M, K, S, P, Q int
	queue            Queue
	links            [154321][]int
	zombie           [154321]int
	dp               [154321]int
	pq               PQ
)

func main() {
	readVariables()
	setZombieDistance()
	dijkstra()
	fmt.Println(dp[N])
	// fmt.Println(zombie[1 : N+1])
	// fmt.Println(dp[1 : N+1])
}

func dijkstra() {
	start := Point{1, 0}
	dp[1] = 0
	heap.Init(&pq)
	heap.Push(&pq, &start)
	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Point)
		// fmt.Println("確定取り出し:", v)
		for _, w := range links[v.loc] {
			if dp[w] != -1 {
				continue
			}
			dp[w] = v.value + cost(w)
			wPoint := Point{w, dp[w]}
			// fmt.Println("確定登録:", wPoint)
			heap.Push(&pq, &wPoint)
		}
	}
}

type Point struct {
	loc   int
	value int
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

func cost(loc int) int {
	INF := 1234567890123456
	if loc == N {
		return 0
	}
	if zombie[loc] == 0 {
		return INF
	}
	if zombie[loc] <= S {
		return Q
	}
	return P
}

func setZombieDistance() {
	for queue.Size() > 0 {
		v := queue.Pop()
		currentDistance := zombie[v]
		for _, w := range links[v] {
			if zombie[w] != -1 {
				continue
			}
			zombie[w] = currentDistance + 1
			queue.Offer(w)
		}
	}
}

func readVariables() {
	N, M, K, S = nextInt(), nextInt(), nextInt(), nextInt()
	P, Q = nextInt(), nextInt()
	//zombie,dp初期化
	for i := 1; i <= N; i++ {
		zombie[i] = -1 //未確定を表す
		dp[i] = -1     //未確定を表す
	}
	for i := 0; i < K; i++ {
		v := nextInt()
		zombie[v] = 0
		queue.Offer(v) //確定済みキュー
	}
	for i := 0; i < M; i++ {
		a, b := nextInt(), nextInt()
		links[a] = append(links[a], b)
		links[b] = append(links[b], a)
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
