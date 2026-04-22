package utils

import "fmt"

func PrintMatrix(matrix [][]int) {
	fmt.Print("[\n")

	for r := range matrix {
		colLen := len(matrix[r])

		for c := range matrix[r] {
			if c == 0 {
				fmt.Print("  ")
			}

			fmt.Printf("(%v %v) = %v", r, c, matrix[r][c])

			if c < colLen-1 {
				fmt.Print(", ")
			} else {
				fmt.Print("\n")
			}
		}
	}

	fmt.Print("]\n")
}
