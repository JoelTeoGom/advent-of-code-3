package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make([][]string, 0)
	symbols := make([]string, 0)

	rows := 0
	var i int
	for scanner.Scan() {
		line := scanner.Text()

		//important
		m = append(m, make([]string, len(line)))
		i = 0
		for col, char := range line {
			m[rows][col] = string(char)
			if (char < '0' || char > '9') && char != '.' {
				symbol := string(char)
				if !contains(symbols, symbol) {
					symbols = append(symbols, string(char))
				}
			}
			i++

		}

		rows++
	}
	fmt.Println(symbols)
	fmt.Println(i)
	fmt.Println(len(m))

	rows = len(m)
	col := i
	sum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			if (m[i][j] < "0" || m[i][j] > "9") && m[i][j] != "." {
				find(&sum, i, j, rows, col, m)
			}
		}
	}

	fmt.Println("Result: ", sum)
}

func find(sum *int, i, j, rows, col int, matrix [][]string) {
	for r := i - 1; r < i+1; r++ {
		for c := j - 1; c < j+1; c++ {
			if (r < rows) && (c < col) {
				fmt.Printf("Matrix[%d][%d] = %s \n", r, c, matrix[r][c])
				if matrix[r][c] > "0" && matrix[r][c] < "9" {
					number, err := strconv.Atoi(matrix[r][c])
					if err != nil {
						fmt.Println("Error :", err)
						return
					}
					*sum += number
				}
			}
		}

	}

}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
