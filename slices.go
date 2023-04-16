package main

import "fmt"

func main() {
	// name := []string{"hesham", "mohamed"}
	// var age []int

	// name = append(name, "Ali", "Naga")
	// age = append(age, 24, 25, 26, 27)

	// fmt.Println(name)
	// fmt.Println(age)

	// destination := make([]string, 2)
	// copy(destination, name)
	// fmt.Println(destination)

	// //pointer on elements 1 and 2 on the name slice
	// nickname := name[1:3]
	// fmt.Println(nickname)

	// name[2] = "Serag"
	// fmt.Println(nickname)

	// fmt.Println(name[1:3])
	// fmt.Println(name[:3])
	// fmt.Println(name[1:])

	//2D slices
	maps := [][]string{
		make([]string, 3),
		make([]string, 5),
		[]string{"a", "b"},
	}
	fmt.Println(maps)
	maps[0] = make([]string, 5)
	maps[1] = make([]string, 7)
	maps[1][6] = "s"
	fmt.Println(maps)
}
