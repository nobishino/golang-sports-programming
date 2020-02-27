package segtree

import (
	"fmt"
	"testing"
)

func ExampleIntSegTree() {
	INF := 1234567890
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	segmentTree := NewIntSegTree(3, min, INF)
	segmentTree.Update(0, 1)
	segmentTree.Update(1, 2)
	segmentTree.Update(2, 3)
	fmt.Println(segmentTree.Find(0, 3))
	fmt.Println(segmentTree.Find(1, 3))
	//Output:
	//1
	//2
}

func TestIntSegtree_Addition(t *testing.T) {
	addition := func(x, y int) int { return x + y }
	segmentTree := NewIntSegTree(3, addition, 0)
	segmentTree.Update(0, 1)
	segmentTree.Update(1, 2)
	segmentTree.Update(2, 3)
	if segmentTree.Find(0, 3) != 6 {
		t.Error()
	}
	if segmentTree.Find(1, 3) != 5 {
		t.Error()
	}
}
