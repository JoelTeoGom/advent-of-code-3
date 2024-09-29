package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make([][]string, 0)

	rows := 0
	var i int
	for scanner.Scan() {
		line := scanner.Text()

		//important
		m = append(m, make([]string, len(line)))
		i = 0
		for col, char := range line {
			m[rows][col] = string(char)
			i++

		}

		rows++
	}

	fmt.Println(i)
	fmt.Println(len(m))

	rows = len(m)
	col := i
	sum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			if (m[i][j] < "0" || m[i][j] > "9") && m[i][j] != "." {
				fmt.Printf("matrix[%d][%d]: %s \n", i, j, m[i][j])
				find(&sum, i, j, rows, col, m)
				fmt.Println("==============================")
			}
		}
	}

	fmt.Println("Result: ", sum)
}

func find(sum *int, i, j, rows, col int, matrix [][]string) {

	for r := i - 1; r <= i+1; r++ {
		for c := j - 1; c <= j+1; c++ {

			if (r < rows) && (c < col) {
				if matrix[r][c] >= "0" && matrix[r][c] <= "9" {
					fmt.Printf("Matrix[%d][%d] = %s \n", r, c, matrix[r][c])
					newr, newc := findNumber(i, j, r, c, rows, col, matrix, sum)
					if newr != -1 && newc == -1 {
						//fmt.Println("SALTO DE LINEA", newr)
						break
					}
					if newr == -1 && newc == -1 {
						return
					}
					newr, newc = r, c
				}
			}
		}

	}

}

func findNumber(symbolX, symbolY, currentX, currentY, maxRows, maxCol int, matrix [][]string, sum *int) (int, int) {

	number := make([]string, 0) // Cambiamos a []string

	// left
	for i := currentY; i >= 0; i-- {
		if matrix[currentX][i] >= "0" && matrix[currentX][i] <= "9" {
			number = append([]string{matrix[currentX][i]}, number...)

		} else {
			break
		}
	}

	// right

	var i int
	for i = currentY + 1; i < maxCol; i++ {
		if matrix[currentX][i] >= "0" && matrix[currentX][i] <= "9" {
			number = append(number, matrix[currentX][i])
		} else {
			break
		}
	}

	//Convert form string[] to string to int
	numberStr := strings.Join(number, "")

	if n, err := strconv.Atoi(numberStr); err == nil {
		fmt.Println("MATRIX : ", n)
		*sum += n

	} else {
		fmt.Println("Error al convertir a número:", err)
	}

	//update current X and Y
	//check if updated X and Y are valid
	currentY = i

	//fmt.Println("currentY : ", currentY)
	currentY = -1
	//fmt.Println("currentY : ", currentY)
	if currentX > symbolX+1 {
		return -1, -1
	}

	return currentX, currentY
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
