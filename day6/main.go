package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	example_map := readExample()
	// input_map := readInput()

	// answer := solution1(example_map)
	// fmt.Println("Answer for example solution 1: ", answer)

	// answer = solution1(input_map)
	// fmt.Println("Answer for input solution 1: ", answer)

	answer := solution2(example_map)
	fmt.Println("Answer for example solution 2: ", answer)

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

	STARTCHAR, POSTIONCHAR := "^", "X"
	x, y := 0, 0

	data, x, y = findStart(data, STARTCHAR, POSTIONCHAR)

	// fmt.Println(x, y)

	data = walkMap(data, x, y)

	// printMap(data)

	// first get all locations of current path minus starting location
	walkedPositions := getWalkedPositions(data, POSTIONCHAR, x, y)

	// fmt.Println(walkedPositions)

	total := 0
	for i := 0; i < len(walkedPositions); i++ {
		fmt.Println(i)

		total += determineCircle(data, walkedPositions[i], x, y)
	}

	// iterate over path and add object to current iteration

	// log step to array one (a1)
	// keep logging
	// if current location is already in a1, stop adding to it, purge all elements before current location in a1 and start adding to a2
	// if add now to second array
	// if a1[index] != a2[index] reset arrays and add path again to a1

	// if a1[0] == a2[0] and a1[1] == a2[1] than we go in circles

	// a1 = 1,2,3,4,5
	// a2 = 1,2

	return total
}

func determineCircle(data [][]string, position []int, x, y int) int {
	data[position[0]][position[1]] = "#"
	walkedRecord := make([][]int, 0)
	BLOCKCHAR, POSTIONCHAR := "#", "X"
	secondRound := false
	nextStep := []int{0, 0}

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
			if secondRound {
				if nextStep[0] == x && nextStep[1] == y {
					return 1
				} else {
					secondRound = false
				}
			} else {
				walkedRecord = append(walkedRecord, []int{x, y})
				isFound, a := findXYInRecord(walkedRecord, x, y)
				if isFound {
					nextStep = a
					secondRound = true
				}
			}

		}
	}

	// walk map but record all positions by adding x and y as "x_Y" to array if x_y already in list
	//walk once more and see if it matches next x_y in the list fi so we walk a circle

	return 0
}

func findXYInRecord(record [][]int, x, y int) (bool, []int) {
	for i := 0; i < len(record)-1; i++ {
		if record[i][0] == x && record[i][1] == y {
			return true, record[i+1]
		}
	}
	return false, nil
}

func getWalkedPositions(data [][]string, positionChar string, startPositionX int, startPositionY int) [][]int {
	positions := make([][]int, 0)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == positionChar && !(i == startPositionX && j == startPositionY) {
				positions = append(positions, []int{i, j})
			}
		}
	}
	return positions
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
