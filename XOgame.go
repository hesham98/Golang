package main

import "fmt"

func main() {
	xoBoard := [3][3]string{}
	var player string
	fmt.Println("which player will begin x or o")

	for {
		if player == "x" || player == "o" {
			fmt.Println("player", player)
			//get inputes ROW AND COLUMNS
			var row int
			fmt.Println("please enter the desired ROW 1 ,2, or 3")
			fmt.Scanln(&row)
			var col int
			fmt.Println("please enter the desired COLUMN 1 ,2, or 3")
			fmt.Scanln(&col)
			//check row and columns is empty or not
			if row <= 3 && col <= 3 {
				if xoBoard[row-1][col-1] != "" {
					fmt.Println("Please enter in empty place")
					continue
				} else {
					xoBoard[row-1][col-1] = player
					fmt.Println(xoBoard[0])
					fmt.Println(xoBoard[1])
					fmt.Println(xoBoard[2])
					//chech if there is ant winner
					for j := 0; j < 3; j++ {

						if (xoBoard[j][0] == player && xoBoard[j][0] == xoBoard[j][1] && xoBoard[j][0] == xoBoard[j][2]) || (xoBoard[0][j] == player && xoBoard[0][j] == xoBoard[1][j] && xoBoard[0][j] == xoBoard[2][j]) {
							fmt.Println("WINNER")
							return
						}

					}
					if (xoBoard[0][0] == player && xoBoard[0][0] == xoBoard[1][1] && xoBoard[0][0] == xoBoard[2][2]) || (xoBoard[0][2] == player && xoBoard[0][2] == xoBoard[1][1] && xoBoard[0][2] == xoBoard[2][0]) {
						fmt.Println("WINNER")
						return
					}

					//switch player
					if player == "x" {
						player = "o"
					} else {
						player = "x"
					}
				}
			} else {
				fmt.Println("Please the right column and row 1 ,2, or 3")
				continue
			}
		} else {
			fmt.Println("Please enter a valid player to start x OR o")
			fmt.Scanln(&player)
		}
	}
}
