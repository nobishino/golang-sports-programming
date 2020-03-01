package lca

import (
	"math"
)

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
		if (t.Depth[v]-t.Depth[w])>>uint64(k)&1 == 1 {
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
