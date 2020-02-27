package segtree

//IntSegTree は、要素型をintに限定したセグメント木です。
type IntSegTree struct {
	size        int
	unitElement int
	operation   func(int, int) int
	depth       int
	leafNum     int
	nodes       []int
}

func NewIntSegTree(size int, operation func(int, int) int, unitElement int) *IntSegTree {
	//要素を格納する木・配列の深さと大きさを計算する
	d, s := 0, 1
	for s < size {
		d++
		s *= 2
	}
	nodes := make([]int, 2*s)
	//要素を初期化する
	for i := range nodes {
		nodes[i] = unitElement
	}
	//値を返す
	return &IntSegTree{
		size:        size,
		unitElement: unitElement,
		operation:   operation,
		depth:       d + 1,
		leafNum:     s,
		nodes:       nodes,
	}
}

/*
Update は、セグメント木の要素の値を更新します。

locationで何番目の要素かを指定し(0-indexed)、valueで更新後の値を指定します。
*/
func (t *IntSegTree) Update(location, value int) {
	//leafの更新
	nodeIndex := t.leafNum - 1 + location
	t.nodes[nodeIndex] = value
	//親ノードの再計算
	for nodeIndex > 0 {
		nodeIndex = (nodeIndex - 1) / 2
		leftChild := t.nodes[2*nodeIndex+1]
		rightChild := t.nodes[2*nodeIndex+2]
		t.nodes[nodeIndex] = t.operation(leftChild, rightChild)
	}
}

//Find returns a_start * a_{start + 1} * ... * a_{end - 1}
//
//i.e. returns reduced value within [start,end)
func (t *IntSegTree) Find(start, end int) int {
	return t.helper(start, end, 0, 0, t.leafNum)
}

func (t *IntSegTree) helper(start, end, nodeIndex, left, right int) int {
	// fmt.Println("args", start, end, nodeIndex, left, right)
	//重複部分なしの場合、単位元を返す
	if right <= start || end <= left {
		return t.unitElement
	}
	//担当範囲がクエリ範囲に完全に含まれる場合、保持している値を返す
	if start <= left && right <= end {
		return t.nodes[nodeIndex]
	}
	//それ以外の場合、子要素にクエリを投げ、マージする
	mid := (left + right) / 2
	return t.operation(
		t.helper(start, end, 2*nodeIndex+1, left, mid),
		t.helper(start, end, 2*nodeIndex+2, mid, right),
	)
}
