package main

import (
	"fmt"
	"math/rand"
	"time"
)

type board [][]int
const row int = 6
const col int = 6
const generations int = 1



func main() {
	gameBoard := New(row, col)
	gameBoard = presetSeed()

	rand.Seed(time.Now().Unix())
	gameBoard.randomSeed()
	gameBoard.Print()
	gameBoard.iteration(generations)
	gameBoard.Print()
}

func presetSeed() b board {
	b = [][]int{
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
	}
	b.Print()
	return b
}


func New(row, col int) board{
	b := make(board, row)
	for i:=0; i<row; i++{
		b[i] = make([]int, col)
	}
	return b
}

func (b board) randomSeed(){
	for i := range(b){
		for j := range(b[i]){
			b[i][j] = rand.Intn(2)
		}
	}

}

func (b board) Print(){
	for i := range(b){
		fmt.Println(b[i])
	}
	fmt.Println()
}

func (mainBoard board) iteration(gen int){
	copyBoard := New(row, col)
	mainBoard.Copy(copyBoard)

	for g:=0; g<gen; g++{
		for i := range(mainBoard){
			for j := range(mainBoard[i]){
				mainBoard[i][j] = copyBoard.gameLogic(i, j)
			}
		}
		mainBoard.Print()
	}
}

func (b board) Copy(copyBoard board){
	for i := range(b){
		for j := range(b[i]){
			copyBoard[i][j] = b[i][j]
		}
	}
}


func (b board) gameLogic(i, j int) (cellStatus int){
	if b[i][j] == 1 {
		switch b.neighbours(i, j) {
			case 0, 1:
				cellStatus = 0
			case 2, 3:
				cellStatus = 1
			default:
				cellStatus = 0
		}
	} else {
		if b.neighbours(i, j) == 3 {
			cellStatus = 1
		}
	}
	return cellStatus
}


func (b board) neighbours(i, j int) int {
	aliveNeighbours := 0
	if b[i][j] == 1 {
		aliveNeighbours = -1
	}
	//fmt.Printf("[%d, %d]  %d  ", i, j, b[i][j])
	x := 0
	iCopy := i
	for r:=0; r<3; r++{
		if i > 0 && i <= len(b){
			for c:=0; c<3; c++{
				if j > 0 && j <= len(b[iCopy]){
					if b[i-1][j-1] == 1{
						aliveNeighbours++
					}
					x++
				}
				j++
			}
			j -= 3
		}
		i++
	}
	//fmt.Printf("  x = %d  Alive Neighbours = %d\n", x, aliveNeighbours)
	return aliveNeighbours
}