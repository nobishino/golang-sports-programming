package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var permutations [][]int = make([][]int, 0, 10)

func main() {
	var maxDepth int
	var p, q []int
	fmt.Scan(&maxDepth)
	p = make([]int, maxDepth)
	q = make([]int, maxDepth)
	for i := 0; i < maxDepth; i++ {
		var x int
		fmt.Scan(&x)
		p[i] = x - 1
	}
	for i := 0; i < maxDepth; i++ {
		var x int
		fmt.Scan(&x)
		q[i] = x - 1
	}

	for i := 0; i < maxDepth; i++ {
		call(i, 0, initialize(maxDepth))
	}
	//なんかアレだけど一応順列は出た
	sorted := make([][]int, len(permutations))
	copy(sorted, permutations)
	/*sort.Slice was introdued Go v1.8 according to godoc*/
	// sort.Slice(sorted, func(i, j int) bool {
	// 	for x := 0; x < len(permutations[i]); x++ {
	// 		if sorted[i][x] != sorted[j][x] {
	// 			return sorted[i][x] < sorted[j][x]
	// 		}
	// 	}
	// 	return false
	// })
	perms := Perms(sorted)
	sort.Sort(perms)

	index := make(map[string]int)
	for i, v := range perms {
		index[toS(v)] = i
	}
	answer := absInt(index[toS(p)] - index[toS(q)])
	fmt.Println(answer)
}

type Perms [][]int

func (p Perms) Len() int {
	return len(p)
}
func (p Perms) Less(i, j int) bool {
	for x := 0; x < len(p[i]); x++ {
		if p[i][x] != p[j][x] {
			return p[i][x] < p[j][x]
		}
	}
	return false
}

func (p Perms) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
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

func toS(perm []int) string {
	chars := make([]string, len(perm))
	for i, v := range perm {
		chars[i] = strconv.Itoa(v)
	}
	// fmt.Println(chars)
	return strings.Join(chars, "")
}
