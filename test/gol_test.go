package main

import (
	//"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	row := 40
	col := 50
	gameboard := new(row, col)
	for i := range gameboard {
		if len(gameboard) != row || len(gameboard[i]) != col {
			t.Error("Expected array of size", row, "x", col)
		}
	}
}

func TestIteration(t *testing.T) {
	gens := 15
	var expectedResult board = [][]int{
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 1, 1, 1, 0},
		[]int{0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0},
	}

	gameBoard := presetSeed()
	gameBoard.iteration(gens)
	if !reflect.DeepEqual(gameBoard, expectedResult) {
		t.Error("The result does not match with what is expected")
	}
}
