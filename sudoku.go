package main

import (
	"fmt"
	"os"
	"strconv"
)

// Check the validity of the users input (os.Args)
func IsValidInput(inputArgsArray []string) bool {
	// Check if there 9 arguments in array
	if len(inputArgsArray) != 9 {
		return false
	}

	for i := 0; i < len(inputArgsArray)-1; i++ {
		// Check if each argument has 9 letters
		if len(inputArgsArray[i]) != 9 {
			return false
		}

		for j := 0; j < len(inputArgsArray[i])-1; j++ {
			// Check if each letter of each argument is a valid character
			if (inputArgsArray[i][j] >= '0' && inputArgsArray[i][j] <= '9') || inputArgsArray[i][j] == '.' {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// Change the string format of the input into two dimentional array of integers
func formatInput(inputArgsArray []string) [9][9]int {
	var boardArr [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if inputArgsArray[i][j] == '.' {
				boardArr[i][j] = 0
			} else {
				boardArr[i][j], _ = strconv.Atoi(string(inputArgsArray[i][j]))
			}
		}
	}
	return boardArr
}

func hasDuplicates(counts [10]int) bool {
	// Check if there are duplicate numbers in Input
	for i := 1; i < len(counts); i++ {
		if counts[i] >= 2 {
			return true
		}
	}
	return false
}

func isBoardValid(board *[9][9]int) bool {
	// Check if there are duplicate numbers in each row
	for row := 0; row < 9; row++ {
		digitCounter := [10]int{}
		// Check if there are duplicate numbers in each column
		for col := 0; col < 9; col++ {
			digitCounter[board[row][col]]++
		}
		if hasDuplicates(digitCounter) {
			return false
		}
	}

	// Check if there are duplicate numbers in each 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			digitCounter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					digitCounter[board[row][col]]++
				}
				if hasDuplicates(digitCounter) {
					return false
				}
			}
		}
	}
	return true
}

// Check if the board has more than zero empty spots (e.g. 0)
func hasEmptySpot(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

// Solves sudoku puzze recursivly**
func PutDigitRec(board *[9][9]int) bool {
	if !hasEmptySpot(board) {
		return true
	}
	for rowIndex := 0; rowIndex < 9; rowIndex++ {
		for colIndex := 0; colIndex < 9; colIndex++ {
			if board[rowIndex][colIndex] == 0 {
				for candidate := 1; candidate <= 9; candidate++ {
					board[rowIndex][colIndex] = candidate
					if isBoardValid(board) {
						if PutDigitRec(board) {
							return true
						}
						board[rowIndex][colIndex] = 0
					} else {
						board[rowIndex][colIndex] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func main() {
	arguments := os.Args[1:]
	isInputValid := IsValidInput(arguments)
	if len(arguments) < 9 {
		// Invalid arguments lenght
		fmt.Println("Error")
	}
	if !isInputValid {
		// Invalid input error
		fmt.Println("Error")
	} else {
		board := formatInput(arguments)
		if !isBoardValid(&board) {
			// The row or column or 3x3 section has more than one same digit
			fmt.Println("Error")
		} else {
			if PutDigitRec(&board) {
				// Print the sudoku board with new line in the end of each row
				for i := 0; i < len(board); i++ {
					for o := 0; o < len(board[i]); o++ {
						if o == 8 {
							fmt.Printf("%v", board[i][o])
						} else {
							fmt.Printf("%v ", board[i][o])
						}
					}
					fmt.Println()
				}
			}
		}
	}
}
