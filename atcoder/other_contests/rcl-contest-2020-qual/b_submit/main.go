package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	N, M   int
	plan   Board
	answer [500]Action
	dx, dy map[string]int
	dirs   [4]string = [4]string{"U", "D", "L", "R"}
)

func init() {
	dx = make(map[string]int)
	dy = make(map[string]int)
	dx["D"] = 1
	dx["U"] = -1
	dy["R"] = 1
	dy["L"] = -1
}

func readVariables() {
	N = nextInt()
	M = nextInt()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			plan[i][j] = nextInt()
		}
	}
}

type Ans []Action

func NewAns() Ans {
	return make([]Action, 500)
}

type Board [50][50]int

func NewBoard() Board {
	b := Board{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			b[i][j] = -1
		}
	}
	b[0][0] = 0
	b[0][N-1] = 1
	b[N-1][0] = 2
	b[N-1][N-1] = 3
	return b
}

func Greedy() (result Ans) {
	result = NewAns()
	b := NewBoard()
	for k := 0; k < M; k++ {
		score := -1234567890
		bestI, bestJ, bestDir := -1, -1, ""
		color := k % 4
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if b[i][j] != color {
					continue
				}
				for _, dir := range dirs {
					s, ok := b.calcDiff(i, j, color, dir)
					if !ok {
						continue
					}
					// fmt.Println(k, i, j, dir, s)
					if score < s {
						score = s
						bestI, bestJ, bestDir = i, j, dir
					}
				}
			}
		}
		b.Update(bestI, bestJ, color, bestDir)
		result[k] = Action{true, bestI, bestJ, bestDir}
	}
	return
}

func (b *Board) calcDiff(i, j, color int, dir string) (score int, ok bool) {
	if b[i][j] != color {
		return
	}
	di := dx[dir]
	dj := dy[dir]
	before, after, fill := 0, 0, 0
	for k := 1; k <= 5; k++ {
		if OutOfRange(i+di*k, j+dj*k) {
			break
		}
		if b[i+di*k][j+dj*k] == -1 {
			fill++
		}
		if b[i+di*k][j+dj*k] == plan[i+di*k][j+dj*k] {
			before++
		}
		if color == plan[i+di*k][j+dj*k] {
			after++
		}
	}
	score = after - before + 2*fill
	ok = true
	return
}

func OutOfRange(i, j int) bool {
	return i < 0 || j < 0 || i > N-1 || j > N-1
}

func (b *Board) Update(i, j, color int, dir string) {
	di := dx[dir]
	dj := dy[dir]
	for k := 1; k <= 5; k++ {
		if OutOfRange(i+di*k, j+dj*k) {
			break
		}
		b[i+di*k][j+dj*k] = color
	}
}

func judge(a Ans) int {
	b := NewBoard()
	for i, v := range a {
		color := i % 4
		if !v.Act {
			continue
		}
		if b[v.I][v.J] != color {
			return -1
		}
		di := dx[v.Dir]
		dj := dy[v.Dir]
		for k := 1; k <= 5; k++ {
			b[v.I+di*k][v.J+dj*k] = color
		}
	}
	return plan.Compare(b)
}

func (c *Board) Compare(b Board) (score int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if c[i][j] == b[i][j] {
				score++
			}
		}
	}
	return score
}

type Action struct {
	Act bool
	I   int
	J   int
	Dir string
}

func main() {
	readVariables()
	solve()
	p(Greedy())
	// fmt.Println(plan)
	// fmt.Println(NewBoard())
	// fmt.Println(NewBoard().Compare(plan))
}
func p(answer Ans) {
	for _, v := range answer {
		if v.Act {
			fmt.Println(v.I, v.J, v.Dir)
		} else {
			fmt.Println(-1)
		}
	}
}

func solve() {
	// for i := 0; i < M; i++ {
	// 	answer[i].Act = true
	// 	answer[i].I = 5 * (i / N)
	// 	answer[i].J = i % N
	// 	answer[i].Dir = "D"
	// }
}

/* 以下、テンプレート*/

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 10000000), 10000000)
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
