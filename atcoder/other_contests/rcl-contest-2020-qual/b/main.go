package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Ans struct {
	Actions []Action
	Score   int
	Board   Board
}

func NewAns() Ans {
	return Ans{make([]Action, 500), 4, NewBoard()}
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
	bests := make([]Ans, 1)
	bests[0] = NewAns()
	for k := 0; k < M; k++ {
		color := k % 4
		saveCount := 1
		nextBests := make([]Ans, saveCount)
		for c := 0; c < len(bests); c++ {
			ans := bests[c]
			b := ans.Board
			var nextActions ScoredActions = make([]ScoredAction, N*N*4)
			aIndex := 0
			for i := 0; i < N; i++ {
				for j := 0; j < N; j++ {
					for _, dir := range dirs {
						s, ok := b.calcDiff(i, j, color, dir)
						var action Action
						if !ok {
							action = Action{false, i, j, dir}
						} else {
							action = Action{true, i, j, dir}
						}
						nextActions[aIndex] = ScoredAction{action, ans.Score + s, c}
						aIndex++
					}
				}
			}
			sort.Sort(nextActions)
			for c := 0; c < saveCount; c++ {
				action := nextActions[c]
				prevState := bests[action.PrevState]
				nextBests[c] = prevState.Append(action.action, action.score, color)
			}
		}
		bests = nextBests
	}
	result = bests[0]
	return
}

func (a Ans) Append(act Action, newScore, color int) (result Ans) {
	result = NewAns()
	result.Actions = a.Actions
	result.Score = newScore
	result.Board = a.Board.Update(act.I, act.J, color, act.Dir)
	return
}

type ScoredAction struct {
	action    Action
	score     int
	PrevState int
}

type ScoredActions []ScoredAction

func (s ScoredActions) Len() int {
	return len(s)
}
func (s ScoredActions) Less(i, j int) bool {
	return s[i].score > s[j].score
}
func (s ScoredActions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (b Board) calcDiff(i, j, color int, dir string) (score int, ok bool) {
	if b[i][j] != color {
		return
	}
	di := dx[dir]
	dj := dy[dir]
	before, after := 0, 0
	for k := 1; k <= 5; k++ {
		if OutOfRange(i+di*k, j+dj*k) {
			break
		}
		if b[i+di*k][j+dj*k] == plan[i+di*k][j+dj*k] {
			before++
		}
		if color == plan[i+di*k][j+dj*k] {
			after++
		}
	}
	score = after - before
	ok = true
	return
}

func OutOfRange(i, j int) bool {
	return i < 0 || j < 0 || i > N-1 || j > N-1
}

func (b Board) Copy() (result Board) {
	result = NewBoard()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			result[i][j] = b[i][j]
		}
	}
	return
}

func (b Board) Update(i, j, color int, dir string) (result Board) {
	result = b.Copy()
	di := dx[dir]
	dj := dy[dir]
	for k := 1; k <= 5; k++ {
		if OutOfRange(i+di*k, j+dj*k) {
			break
		}
		result[i+di*k][j+dj*k] = color
	}
	return
}

func judge(a Ans) int {
	b := NewBoard()
	for i, v := range a.Actions {
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
	p(Greedy())
}
func p(answer Ans) {
	for _, v := range answer.Actions {
		if v.Act {
			fmt.Println(v.I, v.J, v.Dir)
		} else {
			fmt.Println(-1)
		}
	}
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
