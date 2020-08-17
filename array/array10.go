package array

import (
	"fmt"
	"strconv"
)

/*
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

A partially filled sudoku which is valid.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:

Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
Example 2:

Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.
*/

// IsValidSudoku validate sudoku
func IsValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		checkRow := map[int]int{}
		checkColumn := map[int]int{}
		for j := 0; j < 9; j++ {
			// row
			rowByteNumber := board[i][j]
			rowByteToInt, _ := strconv.Atoi(string(rowByteNumber))
			checkRow[rowByteToInt]++

			// column
			columnByteNumber := board[j][i]
			columnByteToInt, _ := strconv.Atoi(string(columnByteNumber))
			checkColumn[columnByteToInt]++
		}

		for i, v := range checkRow {
			if i != 0 && v > 1 {
				return false
			}
		}

		for i, v := range checkColumn {
			if i != 0 && v > 1 {
				return false
			}
		}
	}

	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			checkRow := map[int]int{}
			checkColumn := map[int]int{}
			for row := i; row < i+3; row++ {
				for column := j; column < j+3; column++ {
					rowByteNumber := board[row][column]
					rowByteToInt, _ := strconv.Atoi(string(rowByteNumber))
					checkRow[rowByteToInt]++

					// column
					columnByteNumber := board[column][row]
					columnByteToInt, _ := strconv.Atoi(string(columnByteNumber))
					checkColumn[columnByteToInt]++
				}
				for index, v := range checkRow {
					if index != 0 && v > 1 {
						fmt.Println(checkRow)
						return false
					}
				}

				for index, v := range checkColumn {
					if index != 0 && v > 1 {
						fmt.Println("2")
						return false
					}
				}
			}

		}
	}

	return true
}
