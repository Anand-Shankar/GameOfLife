package main

import "fmt"
//import "pic"

func main() {
	gameboard := [][]int{
		[]int{1, 0, 0, 1, 0},
		[]int{0, 1, 1, 0, 0},
		[]int{0, 0, 0, 1, 0},
		[]int{1, 0, 0, 0, 0},
		[]int{0, 1, 0, 1, 1},
		[]int{1, 1, 1, 1, 1},
	}

	
	copy_gameboard := make([][]int, len(gameboard))
	uint_gameboard := make([][]uint8, len(gameboard))

	
	//[1,4] = = 0,3; 0,4; 1,3 1,4

	//should check (0,3) (0,4) (1,3) (1,4) (2,3 (2,4))

//	var copy_gameboard [][]int


	for p := range gameboard{
			copy_gameboard[p] = make([]int, len(gameboard[p]))
			uint_gameboard[p] = make([]uint8, len(gameboard[p]))
		}

	for generation := 0; generation < 1; generation++ { // for gen < 1001
		for p := range gameboard{
			for q := range (gameboard[p]){
				copy_gameboard[p][q] = gameboard[p][q]
			}
		}

/*	Pic := func(dx, dy int) [][]uint8 {
		for a := range(gameboard){
			for b := range(gameboard[a]){
				if gameboard[a][b] == 1{
					uint_gameboard[a][b] = 255
				} else {
					uint_gameboard[a][b] = 0
				}
			}
		}
		return uint_gameboard
	}
	pic.Show(Pic)*/

	for i := range gameboard {
		for j := range gameboard[i] {
			if gameboard[i][j] == 1 { //could also use original gameboard instead of copy
				switch neighbours(copy_gameboard, i, j) {
					case 0, 1: //cell dies
						gameboard[i][j] = 0
					case 2, 3: //cell lives
						gameboard[i][j] = 1 //Is there a way to maintain status without processing (apart from default)
					default: //dies
						gameboard[i][j] = 0
				}
			} else {
				if neighbours(copy_gameboard, i, j) == 3 {
					gameboard[i][j] = 1 //cell returns to life
				}
			}
		}
	}
		//fmt.Println(gameboard)
	}
	printResult(gameboard)
}

func printResult(gameboard [][]int){
	for i := range(gameboard){
		fmt.Println(gameboard[i])
	}
}



func neighbours(gameboard [][]int, i, j int) int {
	alive_neighbours := 0
	if gameboard[i][j] == 1 {
		alive_neighbours = -1
	}

	/*	if i == 0 {
			i = 1
			row_iter = 1
		} else if i == len(gameboard) {
			i -= 1
			row_iter = 1
		}

		if j == 0 {
			j = 1
			column_iter = 1
		} else if j == len(gameboard[i]) {
			j -= 1
			column_iter = 1
	*/

	/*loop_run := 1
	  if i == 0 || j == 0 || i==len(gameboard) || j == len(gameboard[i]){
	  	loop_run = 0
	  }*/
	fmt.Printf("[%d, %d]  %d  ", i, j, gameboard[i][j])
	x := 0
	i_copy := i
	for row_iter := 0; row_iter < 3; row_iter++ {
		if i > 0 && i <= len(gameboard) {
			for column_iter := 0; column_iter < 3; column_iter++ {
				if j > 0 && j <= len(gameboard[i_copy]) {
					if gameboard[i-1][j-1] == 1 {
						alive_neighbours++
						//fmt.Printf("  A[%d, %d]", i-1, j-1)
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

//--------
 



/*package main

import "fmt"
import "image"
import "image/color"

func main() {
	gameboard := [][]int{
		[]int{1, 0, 0, 1, 0},
		[]int{1, 1, 1, 0, 0},
		[]int{0, 1, 0, 1, 0},
		[]int{1, 0, 0, 0, 0},
		[]int{1, 1, 1, 1, 1},
		//[]int{1, 1, 1, 1, 1},
	}

	uint_gameboard := make([][]uint8, len(gameboard))

	var i, j int

	for i = range gameboard {
		uint_gameboard[i] = make([]uint8, len(gameboard[i]))
		for j = range gameboard[i] {
			if gameboard[i][j] == 1 {
				uint_gameboard[i][j] = 255
			} else if gameboard[i][j] == 0{
				uint_gameboard[i][j] = 0
			}
		}
	}

	border := image.Rect(0, 0, i, j)

	m := image.NewRGBA(border)
	for i := range gameboard {
		for j := range gameboard[i] {
			m.Set(i, j, color.RGBA{uint_gameboard[i][j], 0, 0, 0})
		}
	}

	
	fmt.Println(m)

}
*/