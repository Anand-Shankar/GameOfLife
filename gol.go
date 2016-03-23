package main

import (
	"fmt"
	"math/rand"
	"time"
)

type board [][]int

const (
	row = 6
	col = 6
	generations = 10		// Move closer to usage.
	userSeedPref = 1   //Should I seed the board with random or preset values?
)

func main() {
	gameBoard := createBoard()
	gameBoard.print()
	fmt.Println("............. \n")

	gameBoard.simulate(generations)
}

func createBoard() (newBoard board) {
	if userSeedPref == 1{
		newBoard = presetSeed()
		return
	}
	newBoard = randomBoard()
	return
}

func presetSeed() board {
	return [][]int{
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
	}
}

func randomBoard() board {
	rand.Seed(time.Now().Unix()) //Initializing randomizer
	b := make(board, row)
	for i := range b {
		b[i] = make([]int, col)
		for j := range b[i] {
			b[i][j] = rand.Intn(2)
		}
	}
	return b
}

func (b board) print() {
	for i := range b {
		fmt.Println(b[i])
	}
	fmt.Println()
}

func (b board) simulate(gen int) {
	var copyBoard board
	for g := 0; g < gen; g++ {
		copyBoard = clone(b)
		for i := range b {
			for j := range b[i] {
				b[i][j] = copyBoard.gameLogic(i, j)
			}
		}
		b.print()
	}
}

func clone(b board) board {
	copyBoard := make(board, row)
	for i := range(b){
		copyBoard[i] = make([]int, col)
			for j := range(b[i]){
				copyBoard[i][j] = b[i][j]
			}
	}
	return copyBoard
}

func (b board) gameLogic(i, j int) (cellStatus int) {
	if b[i][j] == 1 {
		switch b.neighbours(i, j) {
		case 0, 1:
			cellStatus = 0
		case 2, 3:
			cellStatus = 1
		default:
			cellStatus = 0 //if neighbours method returns value above 3, cell dies
		}
	} else {
		if b.neighbours(i, j) == 3 {
			cellStatus = 1
		}
	}
	return
}

func (b board) cell(r, c int) int {

	if r >= 0 && c >= 0 && r < len(b) && c < len(b[0]) {
		return b[r][c]
	} 
	return 0
}

func (b board) neighbours(i, j int) int {
	//i, j = 2, 4 //delete

	aliveNeighbours := b.cell(i-1, j-1) + b.cell(i-1, j) + b.cell(i-1, j+1) +
					   b.cell( i , j-1) + 				   b.cell( i , j+1) +
					   b.cell(i+1, j-1) + b.cell(i+1, j) + b.cell(i+1, j+1)

	return aliveNeighbours
}