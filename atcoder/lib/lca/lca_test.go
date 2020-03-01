package lca

import "fmt"

func ExampleLCA() {
	tree := NewLCA(10, 0)
	tree.AddEdge(0, 1)
	tree.AddEdge(0, 2)
	tree.AddEdge(1, 3)
	tree.AddEdge(1, 4)
	tree.AddEdge(4, 5)
	tree.AddEdge(4, 6)
	tree.AddEdge(6, 7)
	tree.AddEdge(2, 8)
	tree.AddEdge(2, 9)
	fmt.Println(tree.Lca(3, 6))
	fmt.Println(tree.Lca(4, 9))
	fmt.Println(tree.Lca(2, 2))
	//Output:
	//1
	//0
	//2
}
