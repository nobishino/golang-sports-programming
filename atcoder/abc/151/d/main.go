package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	x int
	y int
}

var (
	H, W  int
	links map[Node][]Node = make(map[Node][]Node)
	ok    [20][20]bool
	nodes []Node
	dx    = []int{1, 0, -1, 0}
	dy    = []int{0, 1, 0, -1}
)

func main() {
	readVariables()
	var answer int
	for _, from := range nodes {
		for _, to := range nodes {
			answer = MaxInt(answer, shortest(from, to))
		}
	}
	fmt.Println(answer)
}

func adj(n Node) []Node {
	x := n.x
	y := n.y
	var result []Node
	for k := 0; k < 4; k++ {
		v := x + dx[k]
		w := y + dy[k]
		if v < 0 || w < 0 || v >= H || w >= W {
			continue
		}
		if ok[v][w] {
			result = append(result, Node{v, w})
		}
	}
	return result
}

func shortest(from, to Node) int {
	visited := make(map[Node]bool)
	var queue []Node
	dist := make(map[Node]int)
	queue = append(queue, from)
	visited[from] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, w := range adj(v) {
			if visited[w] {
				continue
			}
			dist[w] = dist[v] + 1
			visited[w] = true
			queue = append(queue, w)
		}
	}
	return dist[to]
}

func readVariables() {
	H, W = nextInt(), nextInt()
	for i := 0; i < H; i++ {
		row := nextStr()
		for j := 0; j < W; j++ {
			ok[i][j] = row[j] == '.'
			nodes = append(nodes, Node{i, j})
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
