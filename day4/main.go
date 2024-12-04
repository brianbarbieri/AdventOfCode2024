package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data_example := readExample()
	data_example2 := readExample2()
	data_input := readInput()

	answer := solution1(data_example)
	fmt.Println("Answer for example solution 1: ", answer)

	answer = solution1(data_input)
	fmt.Println("Answer for input solution 1: ", answer)

	answer = solution2(data_example2)
	fmt.Println("Answer for example solution 2: ", answer)

	answer = solution2(data_input)
	fmt.Println("Answer for input solution 2: ", answer)

}

func solution1(data []string) int {
	total := 0

	// xmas horizontally
	for _, element := range data {
		total += strings.Count(element, "XMAS")
	}

	// xmas horizontally rev
	for _, element := range data {
		total += strings.Count(reverse(element), "XMAS")
	}

	// turn array by trans
	datatrans := transpose(data)

	// xmas horizontally by trans
	for _, element := range datatrans {
		total += strings.Count(element, "XMAS")
	}

	// xmas horizontally rev by trans
	for _, element := range datatrans {
		total += strings.Count(reverse(element), "XMAS")
	}

	// diagonal

	var runeGrid [][]rune
	for _, row := range data {
		runeGrid = append(runeGrid, []rune(row))
	}

	// Extract diagonals
	mainDiagonals := getMainDiagonals(runeGrid)

	// xmas horizontally by Diag
	for _, element := range mainDiagonals {
		total += strings.Count(element, "XMAS")
	}

	// xmas horizontally rev by Diag
	for _, element := range mainDiagonals {
		total += strings.Count(reverse(element), "XMAS")
	}

	antiDiagonals := getAntiDiagonals(runeGrid)

	// xmas horizontally by AntiDiag
	for _, element := range antiDiagonals {
		total += strings.Count(element, "XMAS")
	}

	// xmas horizontally rev by AntiDiag
	for _, element := range antiDiagonals {
		total += strings.Count(reverse(element), "XMAS")
	}

	return total
}

func solution2(data []string) int {
	total := 0

	rows := len(data)
	cols := len(data[0])
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			diag := string(data[i-1][j-1]) + string(data[i][j]) + string(data[i+1][j+1])
			antiDiag := string(data[i-1][j+1]) + string(data[i][j]) + string(data[i+1][j-1])

			if (diag == "MAS" && antiDiag == "MAS") || (diag == "SAM" && antiDiag == "MAS") || (diag == "MAS" && antiDiag == "SAM") || (diag == "SAM" && antiDiag == "SAM") {
				total++
			}

		}
	}
	return total
}

func reverse(chars string) string {
	runes := []rune(chars)
	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func transpose(matrix []string) []string {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([]string, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = ""
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j] += string(matrix[i][j])
		}
	}

	return transposed
}

func getMainDiagonals(grid [][]rune) []string {
	n := len(grid)
	var diagonals []string

	// Top-left to bottom-right diagonals (including above main diagonal)
	for colStart := 0; colStart < n; colStart++ {
		var diagonal []rune
		row, col := 0, colStart
		for row < n && col < n {
			diagonal = append(diagonal, grid[row][col])
			row++
			col++
		}
		diagonals = append(diagonals, string(diagonal))
	}

	// Below main diagonal
	for rowStart := 1; rowStart < n; rowStart++ {
		var diagonal []rune
		row, col := rowStart, 0
		for row < n && col < n {
			diagonal = append(diagonal, grid[row][col])
			row++
			col++
		}
		diagonals = append(diagonals, string(diagonal))
	}

	return diagonals
}

func getAntiDiagonals(grid [][]rune) []string {
	n := len(grid)
	var diagonals []string

	// Top-right to bottom-left diagonals (including above anti-diagonal)
	for colStart := n - 1; colStart >= 0; colStart-- {
		var diagonal []rune
		row, col := 0, colStart
		for row < n && col >= 0 {
			diagonal = append(diagonal, grid[row][col])
			row++
			col--
		}
		diagonals = append(diagonals, string(diagonal))
	}

	// Below anti-diagonal
	for rowStart := 1; rowStart < n; rowStart++ {
		var diagonal []rune
		row, col := rowStart, n-1
		for row < n && col >= 0 {
			diagonal = append(diagonal, grid[row][col])
			row++
			col--
		}
		diagonals = append(diagonals, string(diagonal))
	}

	return diagonals
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func openFile(filePath string) []string {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var array []string

	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, line)
	}

	return array
}

func readExample() []string {
	return openFile("day4/example.txt")
}

func readExample2() []string {
	return openFile("day4/example2.txt")
}

func readInput() []string {
	return openFile("day4/input.txt")
}
