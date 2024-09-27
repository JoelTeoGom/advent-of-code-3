package main

import (
	"bufio"
	"fmt"
	"os"
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

	for scanner.Scan() {
		line := scanner.Text()

		m = append(m, make([]string, len(line)))

		for col, char := range line {
			m[rows][col] = string(char)
			if (char < '0' || char > '9') && char != '.' {
				symbol := string(char)
				if !contains(symbols, symbol) {
					symbols = append(symbols, string(char))
				}
			}
		}

		rows++
	}
	fmt.Println(symbols)


	for i

}


func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
