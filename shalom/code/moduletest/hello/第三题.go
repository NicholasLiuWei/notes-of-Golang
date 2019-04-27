package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Chess [6][7]int

type Step struct {
	i, j int
}

//	代表连接一起的相同棋子个数
var times = 4

func (M *Chess) printChess() {
	text := ""
	for _, v := range M {
		for _, chess := range v {
			if chess == 0 {
				text += "."
			} else if chess == 1 {
				text += "x"
			} else if chess == 2 {
				text += "o"
			}
		}
		text += "\n"
	}
	fmt.Println(text)
}

func (M *Chess) moveInChess(x, y, flag int) {
	if x <= len(M) && y <= len(M[1]) {
		M[x][y] = flag
	}
}

func (m *Chess) simulateGame(steps [8]Step) int {
	//	默认黑棋先下
	flag := 1
	for {
		randomLine := rand.Intn(7)

		for i := len(m) - 1; i >= 0; i-- {
			if m[i][randomLine] == 0 {
				m[i][randomLine] = flag
				break
			}
		}
		//	下好一步棋，更换下棋角色
		if flag == 1 {
			flag = 2
		} else if flag == 2 {
			flag = 1
		}
		//	判断是否有玩家胜出
		result, winer := m.isWin(steps)
		if result {
			return winer
		}

		//	判断棋盘是否填满
		if !m.isContinue() {
			break
		}
	}
	return 0
}

func (m *Chess) isContinue() bool {
	for _, v1 := range m {
		for _, v2 := range v1 {
			if v2 == 0 {
				return true
			}
		}
	}
	return false
}

func (m *Chess) isLink(nextI, nextJ, v int, step Step, times int) bool {
	//	time==1时，代表isLink函数已经调用了4次，已经有一方获胜
	if times == 1 {
		return true
	}
	nextI += step.i
	nextJ += step.j
	//	保证不超出棋盘范围
	if (0 <= nextI && nextI < 6) && (0 <= nextJ && nextJ < 7) {
		if m[nextI][nextJ] == v {
			//	继续与相同方向的下一个棋子作比较
			res := m.isLink(nextI, nextJ, v, step, times-1)
			return res
		}
	}
	return false
}

func (m *Chess) isWin(steps [8]Step) (bool, int) {
	//	从下往上遍历,遍历每个点，8个方向是否有相同棋子
	for i := len(m) - 1; i >= 0; i-- {
		for j, v := range m[i] {
			//	如果是空白，则不需要遍历
			if v == 0 {
				continue
			}
			for _, step := range steps {
				nextI := i
				nextJ := j

				result := m.isLink(nextI, nextJ, v, step, times)
				if result {
					return result, v
				}
			}
		}
	}
	return false, 0
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var M Chess
	M.printChess()

	//	8种遍历的可能性
	var steps = [8]Step{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

	//	开始模拟进行游戏
	winer := M.simulateGame(steps)
	if winer == 0 {
		fmt.Println("平局")
	} else if winer == 1 {
		fmt.Println("黑棋获胜！")
	} else if winer == 2 {
		fmt.Println("白棋获胜！")
	}

	M.printChess()
}
