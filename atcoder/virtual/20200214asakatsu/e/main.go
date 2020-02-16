package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	N, M  int
	tasks Tasks
	pq    PriorityQueue
)

type Task struct {
	Delay int
	Value int
}

type Tasks []Task
type PriorityQueue []*Task

func (ts Tasks) Len() int {
	return len(ts)
}

func (ts Tasks) Less(i, j int) bool {
	return ts[i].Delay < ts[j].Delay
}

func (ts Tasks) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (ts PriorityQueue) Len() int {
	return len(ts)
}

func (ts PriorityQueue) Less(i, j int) bool {
	return ts[i].Value > ts[j].Value
}

func (ts PriorityQueue) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	task := x.(*Task)
	*pq = append(*pq, task)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return x
}

func readVariables() {
	N, M = nextInt(), nextInt()
	tasks = make(Tasks, N)
	pq := make(PriorityQueue, 0)
	for i := 0; i < N; i++ {
		d, v := nextInt(), nextInt()
		tasks[i] = Task{d, v}
	}
	heap.Init(&pq)
}

func main() {
	readVariables()
	sort.Sort(tasks)
	answer := 0
	i := 0
	j := 1
	for {
		if j > M {
			break
		}
		for i < N && tasks[i].Delay <= j {
			heap.Push(&pq, &tasks[i])
			i++
		}
		// fmt.Println(pq.Len())
		if pq.Len() > 0 {
			t := heap.Pop(&pq).(*Task)
			answer += t.Value
		}
		// fmt.Println(j, answer)
		j++
	}
	fmt.Println(answer)
}

/* 以下、テンプレート*/

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1000000), 1000000)
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
