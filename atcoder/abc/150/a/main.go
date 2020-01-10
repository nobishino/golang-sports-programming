package main

import (
	"fmt"
)

var (
	a, b int
)

func main() {
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a*500 >= b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
