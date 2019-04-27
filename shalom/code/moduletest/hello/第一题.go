package main

import (
	"fmt"
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

func main() {
	var M Chess
	M.printChess()

	//	传入横纵坐标和下棋角色
	M.moveInChess(0, 3, 1)
	M.moveInChess(3, 2, 1)
	M.moveInChess(2, 3, 2)
	M.moveInChess(3, 5, 2)

	M.printChess()
}
