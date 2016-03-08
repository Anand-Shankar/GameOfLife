package main

import "fmt"

func main() {
	gameboard := [][]int{
		[]int{1, 0, 0, 1, 0},
		[]int{0, 1, 1, 0, 0},
		[]int{0, 0, 0, 1, 0},
		[]int{1, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1},
		[]int{1, 1, 1, 1, 1},
	}

	for generation := 0; generation < 3; generation++ { // for gen < 1001
		copy_gameboard := gameboard

		for i, _ := range gameboard {
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
		fmt.Println(gameboard)
	}
}

func neighbours(gameboard [][]int, i, j int) int {
	alive_neighbours := 0
	

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

loop_run := 1
if i == 0 || j == 0 || i==len(gameboard) || j == len(gameboard[i]){
	loop_run = 0
}
	}
	x:=1
	for row_iter := 0; row_iter < 3; row_iter++ {
		if i!=0 && i!= len(gameboard){
			for column_iter	:= 0 ; column_iter < 3; column_iter++ {
				if j !=0 && j != len(gameboard){
					if gameboard[i-1][j-1] == 1 {
						alive_neighbours++
					}			
				}	
				x+=1
				q++		
			}
		}
		i++
	}
	fmt.Println(x)
	return alive_neighbours
}




//--------


