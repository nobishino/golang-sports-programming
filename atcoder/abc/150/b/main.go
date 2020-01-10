package main

import (
	"fmt"
)

var (
	n int
	s string
)

func main() {
	fmt.Scan(&n)
	fmt.Scan(&s)
	var answer int
	for i := 0; i < n-2; i++ {
		if s[i] == 'A' && s[i+1] == 'B' && s[i+2] == 'C' {
			answer++
		}
	}
	fmt.Println(answer)
}
