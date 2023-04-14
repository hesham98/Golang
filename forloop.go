package main

import "fmt"

func main() {
	// for i := 0; i < 4; i++ {
	// 	if i == 1 {
	// 		break
	// 	} else {
	// 		fmt.Println("hi", i)
	// 	}
	// }

	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		if i == 1 && j == 1 {
	// 			break
	// 		} else {
	// 			fmt.Println("x =", i, "y =", j)
	// 		}
	// 	}
	// }

	b := false
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 0 {
				b = true
				break
			}
			fmt.Println("x =", i, "y =", j)
		}
		if b == true {
			break
		}
	}
	return
}
