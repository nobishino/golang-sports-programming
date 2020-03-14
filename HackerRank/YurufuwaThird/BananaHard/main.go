package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	N, Q  int
	links [154321][]Link
	MOD   int = 1000000007
	tree  RootedTree
	vals  []int
)

type Link struct {
	dest int
	m    int
}

func readVariables() {
	N = nextInt()
	tree = NewLCA(N+1, 0)
	for i := 0; i < N-1; i++ {
		u, v, c := nextInt(), nextInt(), nextInt()
		links[u] = append(links[u], Link{v, c})
		links[v] = append(links[v], Link{u, c})
		tree.AddEdge(u, v)
	}
	Q = nextInt()
}

func main() {
	readVariables()
	pre(1)
	for q := 0; q < Q; q++ {
		m, p, x := nextInt(), nextInt(), nextInt()
		a := query(m, p, x)
		fmt.Println(a)
	}
}

func pre(m int) {
	vals = make([]int, N+1)
	for i := range vals {
		vals[i] = -1
	}
	vals[m] = 1
	var q Queue = make([]int, 0)
	q.Offer(m)
	for len(q) > 0 {
		v := q.Pop()
		for _, link := range links[v] {
			w := link.dest
			c := link.m
			if vals[w] != -1 {
				continue
			}
			vals[w] = c * vals[v]
			vals[w] %= MOD
			q.Offer(w)
		}
	}
}

func query(m, p, x int) int {
	result := (vals[m] * vals[p]) % MOD
	lca := tree.Lca(m, p)
	inv := ModPow(vals[lca], MOD-2, MOD)
	inv = (inv * inv) % MOD
	result *= (inv * x) % MOD
	result %= MOD
	return result
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

type Queue []int

func (q *Queue) Offer(x int) {
	*q = append(*q, x)
}
func (q *Queue) Pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

type Tree struct {
	Size int
	Adj  [][]int
}

type RootedTree struct {
	Tree
	Root        int
	Depth       []int
	Parent      [][]int //Parent[v][k] = 2^k th parent of v
	maxLogSize  int
	initialized bool
}

/*
NewLCA returns rooted tree of size n.
*/
func NewLCA(size, root int) RootedTree {
	tree := Tree{
		Size: size,
		Adj:  make([][]int, size),
	}
	MaxLogSize := int(math.Ceil(math.Log2(float64(size)))) + 1
	rootedTree := RootedTree{
		Tree:       tree,
		Root:       root,
		Depth:      make([]int, size),
		Parent:     make([][]int, size),
		maxLogSize: MaxLogSize,
	}
	return rootedTree
}

/*
Init prepares for incoming Lca() queries.
*/
func (t *RootedTree) Init() {
	for i := 0; i < t.Size; i++ {
		t.Parent[i] = make([]int, t.maxLogSize)
	}
	t.dfs(t.Root, -1, 0)
	for k := 0; k < t.maxLogSize-1; k++ {
		for v := 0; v < t.Size; v++ {
			if t.Parent[v][k] == -1 {
				t.Parent[v][k+1] = -1
			} else {
				t.Parent[v][k+1] = t.Parent[t.Parent[v][k]][k]
			}
		}
	}
	t.initialized = true
}

func (t *RootedTree) dfs(v, parent, depth int) {
	t.Parent[v][0] = parent
	t.Depth[v] = depth
	for _, w := range t.Adj[v] {
		if w == parent {
			continue
		}
		t.dfs(w, v, depth+1)
	}
}

/*
AddEdge adds new edge connecting the v-th and the w-th node to the rooted tree instance.
*/
func (t *RootedTree) AddEdge(v, w int) {
	t.Adj[v] = append(t.Adj[v], w)
	t.Adj[w] = append(t.Adj[w], v)
}

/*
Lca returns the lowest common ancestor of the v-th and w-th node.
*/
func (t *RootedTree) Lca(v, w int) int {
	if !t.initialized {
		t.Init()
	}
	if t.Depth[v] < t.Depth[w] {
		v, w = w, v //ensure v >= w, v is not shallower than w
	}
	for k := 0; k < t.maxLogSize; k++ {
		if (t.Depth[v]-t.Depth[w])>>k&1 == 1 {
			v = t.Parent[v][k]
		}
	}
	if v == w {
		return v
	}
	for k := t.maxLogSize - 1; k >= 0; k-- {
		if t.Parent[v][k] != t.Parent[w][k] {
			v = t.Parent[v][k]
			w = t.Parent[w][k]
		}
	}
	return t.Parent[v][0]
}
