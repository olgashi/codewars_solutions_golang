package main

import (
	"fmt"
)


func ValidateSolution(m [][]int) bool {
	for row := 0; row < 9; row++ {
		uniqueDigitsListCol := [10]int{}
		uniqueDigitsListRow := [10]int{}
		for col := 0; col < 9; col++ {

			rowCell := m[row][col]
			colCell := m[col][row]

			if (rowCell == 0 || colCell == 0) || uniqueDigitsListRow[rowCell] == 1 || uniqueDigitsListCol[colCell] == 1 {
				return false
			} 
			uniqueDigitsListRow[rowCell], uniqueDigitsListCol[colCell] = 1, 1			
		}
	}
  lSqRow := 0
	lSqCol := 0

	littleSquareNums :=[10]int{}

	for lSqCol < 9 {
		for i := 0; i < 3; i++ {
			cell := m[lSqRow][lSqCol]
			if (littleSquareNums[cell] == 1) {
				return false
			}
			littleSquareNums[cell] = 1

			lSqCol++
		}
		lSqRow++
		lSqCol = lSqCol - 3

		if lSqRow % 3 == 0 {
	    littleSquareNums = [10]int{}
		}

		if lSqRow == 9 {
			lSqRow = 0
			lSqCol += 3
		}
	}

return true
}


func main() {
	fmt.Println(ValidateSolution([][]int{
		{5, 3, 4, 6, 7, 8, 9, 2, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
  }))

}