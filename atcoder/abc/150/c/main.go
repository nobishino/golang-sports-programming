package main

import (
	"fmt"
	"sort"
)

//スクロールされると動く
//permutations 意味　順列
//欠陥では...
//フォローしなければOK
//フォローってなんだっけ
//とりあえずファイルエクスプローラから別なファイル開けばフォローされないような気がする　知らんけど
var permutations [][]int = make([][]int, 0, 10)

//これ今ほかに誰かいるのかな only me
//See -> Session Details > Participants
//make sense, well noted
func main() {
	fmt.Println("Hello")
	maxDepth := 4
	for i := 0; i < maxDepth; i++ {
		call(i, 0, initialize(maxDepth))
	}
	//なんかアレだけど一応順列は出た
	//辞書順ソートするか
	fmt.Println(permutations)
	fmt.Println(len(permutations))
	sorted := make([][]int, len(permutations))
	copy(sorted, permutations)
	sort.Slice(sorted, func(i, j int) bool {
		c++
		for x := 0; x < len(permutations[i]); x++ {
			if sorted[i][x] != sorted[j][x] {
				return sorted[i][x] < sorted[j][x]
			}
		}
		return true
	})
	for i := 0; i < len(permutations); i++ {
		fmt.Println(permutations[i], sorted[i])
	}
	fmt.Println(c)
}

var c int

func initialize(length int) (result []int) {
	result = make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = -1
	}
	return
}

//DFSみたいなので順列を求めようとしています
func call(v int, depth int, path []int) {
	copiedPath := make([]int, len(path))
	maxDepth := len(path)
	copy(copiedPath, path)
	copiedPath[v] = depth
	// fmt.Println(copiedPath)
	if depth == maxDepth-1 {
		permutations = append(permutations, copiedPath)
	}
	for i, w := range copiedPath {
		if w != -1 {
			continue
		} else {
			call(i, depth+1, copiedPath)
		}
	}
}
