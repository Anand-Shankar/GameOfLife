package main

import "fmt"

//import "pic"

func main() {
	gameboard := [][]int{
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 1, 1, 1, 0},
		[]int{0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
	}

	copy_gameboard := make([][]int, len(gameboard))
	uint_gameboard := make([][]uint8, len(gameboard))

	for p := range gameboard {
		copy_gameboard[p] = make([]int, len(gameboard[p]))
		uint_gameboard[p] = make([]uint8, len(gameboard[p]))
	}

	for generation := 0; generation < 5; generation++ {
		for p := range gameboard {
			for q := range gameboard[p] {
				copy_gameboard[p][q] = gameboard[p][q]
			}
		}

		for i := range gameboard {
			for j := range gameboard[i] {
				if gameboard[i][j] == 1 {
					switch neighbours(copy_gameboard, i, j) {
					case 0, 1:
						gameboard[i][j] = 0
					case 2, 3:
						gameboard[i][j] = 1
					default:
						gameboard[i][j] = 0
					}
				} else {
					if neighbours(copy_gameboard, i, j) == 3 {
						gameboard[i][j] = 1
					}
				}
			}
		}
		printResult(gameboard)
	}

}

func printResult(gameboard [][]int) {
	for i := range gameboard {
		fmt.Println(gameboard[i])
	}
}

func neighbours(gameboard [][]int, i, j int) int {
	alive_neighbours := 0
	if gameboard[i][j] == 1 {
		alive_neighbours = -1
	}

	fmt.Printf("[%d, %d]  %d  ", i, j, gameboard[i][j])
	x := 0
	i_copy := i
	for row_iter := 0; row_iter < 3; row_iter++ {
		if i > 0 && i <= len(gameboard) {
			for column_iter := 0; column_iter < 3; column_iter++ {
				if j > 0 && j <= len(gameboard[i_copy]) {
					if gameboard[i-1][j-1] == 1 {
						alive_neighbours++

					}
					x++
				}
				j++
			}
			j -= 3
		}
		i++
	}
	fmt.Printf("  x = %d  Alive Neighbours = %d\n", x, alive_neighbours)
	return alive_neighbours
}



func newGameboard(row, column int) [][]int{

}

type board [][]int


gameBoard := board{}

const row, column := 5, 5
gameBoard.newGameboard(row, column)

gameBoard.seed(r,c)



gameBoard.randomSeed()
gameBoard.iteration()
gameBoard.run(iterCount)

assert



/*
package main

import "fmt"

func main() {

	type board [][]int
	
	gameBoard := board{}
	gameBoard = [][]int{
		[]int{5, 5},
		[]int{6, 6},
	}
	
	
	//gameBoard.newGameboard(row, column)

//fmt.Println(game.board)
//fmt.Println(game)
fmt.Println(gameboard)

}


type board [][]int
gameBoard := board{}

const row = 5
const column = 5


func newGameboard(row, column int) {
	gameBoard := make(int[][], len(i))
	for i:=0; i<row; i++{
		make(int[], len(gameBoard[i]))
		for j:=0; j<column; j++{
			gameBoard*/


/*			type board [][]int
gameBoard := board{}

const row = 5
const column = 5


func (g gameBoard) newGameboard(row, column int) {
	gameBoard := make(int[][], row)
	for i:=0; i<row; i++{
		make(int[], column)
	}
}
*/