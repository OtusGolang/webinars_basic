package main

import "fmt"

func main() {
	var grid string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				grid += " "
			} else {
				grid += "#"
			}
		}
		grid += "\n"
	}
	fmt.Println(grid)
}
