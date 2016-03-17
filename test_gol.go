package main

import (
	"fmt"
	"math/rand"
)


type board [][]int
const row int = 5
const col int = 5

func New(row, col int) board{
	b := make(board, row)
	for i:=0; i<row; i++{
		b[i] = make([]int, col)
	}
	fmt.Println(len(b))
	return b
}

func (g gameBoard) randomSeed(){
	for i := range(g){
		for j := range(g[i]){
			g[i][j] = 
		}
	}

}


func main() {
	gameBoard := New(row, col)
	gameBoard.randomSeed()
	fmt.Printf("%T\n", gameBoard)
	fmt.Println(len(gameBoard))
	fmt.Println(gameBoard)


	fmt.Println(rand.Intn(1))

}

/*package main

import "fmt"
import "math/rand"
import "time"

func main() {
	rand.Seed(time.Now().Unix())
	randomNumber := rand.Intn(134254)
	fmt.Println(randomNumber)
}*/