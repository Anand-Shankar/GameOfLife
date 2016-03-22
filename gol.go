package main

import (
	"fmt"
	"math/rand"
	"time"
)

type board [][]int

// Collapse into one `const` block.
// No need to give type to a const.
const (
	row = 6
	col = 6
	generations = 1 // Move closer to usage.
	userSeedPref = 1   //Should I seed the board with random or preset values?
)

func main() {
	// Don't have commented code.
	
	//gameBoard := presetSeed()
	gameBoard := createBoard()

	gameBoard.print()
	fmt.Println("............. \n")
	gameBoard.simulate(generations)
	gameBoard.print()
}


func createBoard() (newBoard board) {
	//fmt.Scanln(userSeedPref)
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
		// Ideally you dont want to mix orthogonal concerns, like 
		// simulating a board vs printing the output - which in the near
		// future could be to image file, or HTML, etc - in the same function.
		// That breaks SRP. Therefore, explore other ways of recording history.
		// One way would be: return a slice of boards which are essentially the
		// generations as the simulation ran. Then the calling method is free to
		// do with it as it pleases.
	}
}

// Idiomatically, this should be called `clone`.
// Read about deep copy.
func clone(b board) board {
	copyBoard := make(board, row)
	for i := range(b){
		copyBoard[i] = make([]int, col)
	}
	copy(copyBoard, b)
	return copyBoard
}

// No need for a return variable name. You never use it apart from a indirect return.
// Be explicit. A lot of opportunies for a quick and early return.
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
	return cellStatus
}

func (b board) neighbours(i, j int) int {

	aliveNeighbours := 0
	if b[i][j] == 1 {
		aliveNeighbours = -1
	}
	// Explore using a log package when you can output debug lines.
	//fmt.Printf("[%d, %d]  %d  ", i, j, b[i][j])
	
	// See if there is a better way of calculating neighbours.
	// Ex: return b.cell(i - 1, j - 1) + b.cell(i - 1, j) 
	
	x := 0
	iCopy := i
	for r := 0; r < 3; r++ {
		if i > 0 && i <= len(b) {
			for c := 0; c < 3; c++ {
				if j > 0 && j <= len(b[iCopy]) {
					if b[i-1][j-1] == 1 {
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