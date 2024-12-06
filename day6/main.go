package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	example_map := readExample()
	input_map := readInput()

	answer := solution1(example_map)
	fmt.Println("Answer for example solution 1: ", answer)

	answer = solution1(input_map)
	fmt.Println("Answer for input solution 1: ", answer)

}

func solution1(data [][]string) int {
	STARTCHAR, POSTIONCHAR := "^", "X"
	x, y := 0, 0

	data, x, y = findStart(data, STARTCHAR, POSTIONCHAR)

	data = walkMap(data, x, y)

	total := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == POSTIONCHAR {
				total++
			}
		}
	}
	return total
}

func solution2(data [][]string) int {

	// first get all locations of current path minus starting location
	// iterate over path and add object to current iteration

	// log step to array one (a1)
	// keep logging
	// if current location is already in a1, stop adding to it, purge all elements before current location in a1 and start adding to a2
	// if add now to second array
	// if a1[index] != a2[index] reset arrays and add path again to a1

	// a1 = 1,2,3,4,5
	// a2 = 1,2

	return 0
}

func findStart(data [][]string, startChar string, positionChar string) ([][]string, int, int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == startChar {
				data[i][j] = positionChar
				return data, i, j
			}
		}
	}
	return data, 0, 0
}

func walkMap(data [][]string, x, y int) [][]string {
	BLOCKCHAR, POSTIONCHAR := "#", "X"

	direction := "up"
	for {
		switch direction {
		case "up":
			x--
		case "down":
			x++
		case "left":
			y--
		case "right":
			y++
		}
		if isOutside(data, x, y) {
			break
		}
		if data[x][y] == BLOCKCHAR {
			switch direction {
			case "up":
				x++
				direction = "right"
			case "down":
				x--
				direction = "left"
			case "left":
				y++
				direction = "up"
			case "right":
				y--
				direction = "down"
			}
		} else {
			data[x][y] = POSTIONCHAR
		}
	}
	return data
}

func printMap(data [][]string) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Print(data[i][j])
		}
		fmt.Println()
	}
	fmt.Println(" ")
}

func isOutside(data [][]string, x, y int) bool {
	return x < 0 || x >= len(data) || y < 0 || y >= len(data[0])
}

func openFile(filePath string) [][]string {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var array [][]string

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		array = append(array, splitLine)
	}

	return array
}

func readExample() [][]string {
	return openFile("day6/example.txt")
}

func readInput() [][]string {
	return openFile("day6/input.txt")
}
