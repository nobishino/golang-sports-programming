package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	s   string
	cnt map[byte]int = make(map[byte]int)
)

func main() {
	readVariables()
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}
	var ok bool
	for v := 0; v < 1e4; v += 8 {
		target := strconv.Itoa(v)
		if len(target) < 4 && len(target) < len(s) {
			continue
		}
		c := make(map[byte]int)
		for i := 0; i < len(target); i++ {
			c[target[i]]++
		}
		var ng bool
		for k, v := range c {
			if cnt[k] < v {
				ng = true
			}
		}
		if !ng {
			ok = true
			break
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readVariables() {
	s = nextStr()
}

/* Template */

var scanner *bufio.Scanner

func init() {
	Max := 1001001
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, Max), Max)
	scanner.Split(bufio.ScanWords)
}

func nextStr() string {
	if !scanner.Scan() {
		panic("No more token.")
	}
	return scanner.Text()
}
