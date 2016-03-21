package main

import (
	"fmt"
	"math/rand"
	"time"
)

type board [][]int

// Collapse into one `const` block.
// No need to give type to a const.
const row int = 6
const col int = 6
// Move closer to usage.
const generations int = 10

func main() {
	// Don't have commented code.
	//gameBoard := New(row, col)

	gameBoard := presetSeed()

	// Do this in a init() or right at the start of main()
	rand.Seed(time.Now().Unix())
	//gameBoard.randomSeed()

	gameBoard.Print()
	gameBoard.iteration(generations)
	// We are printing too many times. The inner loop inside
	// `iteration` also is printing it at the end.
	gameBoard.Print()
}

// Why it is not simply returning the value?
func presetSeed() (b board) {
	return [][]int{
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
	}
}

func new(row, col int) board {
	b := make(board, row)
	for i := 0; i < row; i++ {
		b[i] = make([]int, col)
	}
	return b
}

// Can this be a `randomBoard` function which just returns
// a randomized board.
func (b board) randomSeed() {
	for i := range b {
		for j := range b[i] {
			b[i][j] = rand.Intn(2)
		}
	}
}

// Why is the method exported?
func (b board) Print() {
	for i := range b {
		fmt.Println(b[i])
	}
	fmt.Println()
}

// Have a consistent short name for the method receiver.
// Also, is `run` / `simulate` a better name?
func (mainBoard board) iteration(gen int) {
	copyBoard := new(row, col)
	mainBoard.Copy(copyBoard)

	for g := 0; g < gen; g++ {
		mainBoard.Copy(copyBoard)
		for i := range mainBoard {
			for j := range mainBoard[i] {
				mainBoard[i][j] = copyBoard.gameLogic(i, j)
			}
		}
		// Ideally you dont want to mix orthogonal concerns, like 
		// simulating a board vs printing the output - which in the near
		// future could be to image file, or HTML, etc - in the same function.
		// That breaks SRP. Therefore, explore other ways of recording history.
		// One way would be: return a slice of boards which are essentially the
		// generations as the simulation ran. Then the calling method is free to
		// do with it as it pleases.
		mainBoard.Print()
	}
}

// Idiomatically, this should be called `clone`.
// Read about deep copy.
func (b board) Copy(copyBoard board) {
	for i := range b {
		for j := range b[i] {
			copyBoard[i][j] = b[i][j]
		}
	}
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
			// Avoid a return 0 twice.
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
	// Leave a comment explaining behavior when it cannot be deduced in under 5 seconds.
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
