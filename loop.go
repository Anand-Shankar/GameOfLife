package main

import "fmt"

var x int
var i int
var j int

func main() {

	x = 0
	i = 0
	j = 0

	for ; i < 3; i++ {
		for q := j; q < 3; q++ {
			x++
		}
		q++
	}
	fmt.Println(x, i, j)
}
