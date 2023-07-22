package main

import "fmt"

func main() {
	var size int

	fmt.Scanf("%d", &size)

	var grid string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
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
