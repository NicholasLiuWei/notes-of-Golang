package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Chess [6][7]int

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



func (m *Chess)simulateGame()  {
	flag := 1
	for{
		randomLine := rand.Intn(7)

		for i := len(m) - 1; i >= 0; i-- {
			//	判断是否有空白
			if m[i][randomLine] == 0 {
				m[i][randomLine] = flag
				break
			}
		}
		//	更换下棋角色
		if flag==1{
			flag = 2
		}else if flag ==2{
			flag = 1
		}
		//	判断棋盘是否填满
		if !m.isContinue(){
			break
		}
	}
}

func (m *Chess)isContinue() bool{
	for _, v1 := range m {
		for _, v2 := range v1 {
			if v2 == 0{
				return true
			}
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var M Chess
	M.printChess()

	//	开始模拟进行游戏
	M.simulateGame()

	M.printChess()
}