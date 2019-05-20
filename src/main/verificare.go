package main

import "fmt"

func main() {
	sudoku := [9][9]uint{
		{6, 8, 2, 1, 9, 4, 3, 5, 7},
		{7, 3, 1, 5, 6, 8, 9, 2, 4},
		{4, 9, 5, 7, 2, 3, 8, 6, 1},
		{8, 2, 7, 9, 3, 5, 1, 4, 6},
		{5, 1, 9, 6, 4, 7, 2, 8, 3},
		{3, 6, 4, 2, 8, 1, 5, 7, 9},
		{9, 5, 6, 4, 1, 2, 7, 3, 8},
		{2, 4, 8, 3, 7, 9, 6, 1, 5},
		{1, 7, 3, 8, 5, 6, 4, 9, 1},
	}
	if check(sudoku) {
		fmt.Println("Rezolvare corecta!")
	}
}

func check(sudoku [9][9]uint) bool {
	var a = [3][9]uint{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row := a[0][i]
			column := a[1][j]
			square := a[2][3*(i/3)+(j/3)]
			position := sudoku[i][j]
			if (row & (1 << position)) == (1 << position) {
				fmt.Println("Greseala pe randul ", i + 1)
				return false
			} else {
				a[0][i] = row | (1 << position)
			}
			if (column & (1 << position)) == (1 << position) {
				fmt.Println("Greseala pe coloana ", j + 1)
				return false
			} else {
				a[1][j] = column | (1 << position)
			}
			if (square & (1 << position)) == (1 << position) {
				fmt.Println("Greseala in patrat")
				return false
			} else {
				a[2][3*(i/3)+(j/3)] = square | (1 << position)
			}
		}
	}
	return true
}