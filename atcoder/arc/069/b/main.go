package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N        int
	S        string
	patterns [4]string = [4]string{"SS", "SW", "WS", "WW"}
	sheep    []bool
	strings  []bool
)

func readVariables() {
	N = nextInt()
	S = nextStr()
	// fmt.Scan(&N)
	// fmt.Scan(&S)
	sheep = make([]bool, N)
	strings = make([]bool, N)
	// fmt.Println(N, len(strings), len(S))
	// fmt.Println("S=", S)
	// fmt.Println(sheep, strings)
	for i := 0; i < N; i++ {
		c := string(S[i])
		strings[i] = c == "o"
	}
}

func main() {
	readVariables()
	var c int
	// fmt.Println(strings)
	for _, v := range patterns {
		sheep[0] = string(v[0]) == "S"
		sheep[N-1] = string(v[1]) == "S"
		sheep[1] = g(sheep[N-1], sheep[0], strings[0])
		// fmt.Println(sheep[1], g(sheep[N-1], sheep[0], strings[0]), sheep[N-1], sheep[0], strings[0])
		for i := 2; i < N-1; i++ {
			sheep[i] = g(sheep[i-2], sheep[i-1], strings[i-1])
		}
		ok1 := sheep[N-1] == g(sheep[N-3], sheep[N-2], strings[N-2])
		ok2 := sheep[0] == g(sheep[N-2], sheep[N-1], strings[N-1])
		if ok1 && ok2 {
			// fmt.Println(sheep)
			break
		}
		c++
	}
	// fmt.Println(c)
	if c < 4 {
		for _, v := range sheep {
			if v {
				fmt.Print("S")
			} else {
				fmt.Print("W")
			}
		}
		fmt.Println()
	} else {
		fmt.Println(-1)
	}
}

func g(left, center, saysSame bool) bool {
	same := center && saysSame || !center && !saysSame
	if same {
		return left
	} else {
		return !left
	}
}

/* 以下、テンプレート*/

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1000000), 1000000)
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
