package main

import "fmt"

func main() {

	// name := [4]string{"hesham", "mohamed", "ali", "naga"}
	// fmt.Println(name)
	// name[0] = "mohamed"
	// name[1] = "ali"
	// name[2] = "naga"
	// name[3] = "serag"
	// fmt.Println(name)

	// names := [3]string{}
	// fmt.Println(names)
	// arraySize := len(names)
	// fmt.Println(arraySize)

	XOBoard := [3][3]string{}
	XOBoard = [3][3]string{
		[3]string{"-", "-", "x"},
		[3]string{"-", "o", "-"},
		[3]string{"x", "o", "-"},
	}
	XOBoard[0][1] = "o"
	fmt.Println(XOBoard[0])
	fmt.Println(XOBoard[1])
	fmt.Println(XOBoard[2])
}
