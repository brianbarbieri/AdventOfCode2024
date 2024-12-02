package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// data_example := readExample()
	data_input := readInput()

	// answer := solution1(data_example)
	// fmt.Println("Answer for example solution 1: ", answer)

	// answer := solution1(data_input)
	// fmt.Println("Answer for input solution 1: ", answer)

	// answer = solution2(data_example)
	// fmt.Println("Answer for example solution 2: ", answer)

	answer := solution2(data_input)
	fmt.Println("Answer for input solution 2: ", answer)

}

func solution1(data [][]int) int {
	total := 0

	for i := 0; i < len(data); i++ {
		isNegative := (data[i][0] - data[i][1]) < 0
		valid := true // Track if the row is valid

		for j := 0; j < len(data[i])-1; j++ {
			diff := data[i][j] - data[i][j+1]

			// Check each condition
			if diff == 0 || math.Abs(float64(diff)) > 3 || (diff > 0 && isNegative) || (diff < 0 && !isNegative) {
				valid = false // Mark row as invalid
				break
			}
		}

		// Increment total only if the row is valid
		if valid {
			total++
		}
	}

	return total
}

func solution2(data [][]int) int {
	total := 0
	for _, report := range data {
		if checkCondition(report) && dropLevelCheck(report) {
			total++
		}
	}
	return total
}

func checkCondition(inputList []int) bool {
	sortedAsc := make([]int, len(inputList))
	sortedDesc := make([]int, len(inputList))
	copy(sortedAsc, inputList)
	copy(sortedDesc, inputList)

	sort.Ints(sortedAsc)
	sort.Slice(sortedDesc, func(i, j int) bool {
		return sortedDesc[i] > sortedDesc[j]
	})

	if !equal(inputList, sortedAsc) && !equal(inputList, sortedDesc) {
		return false
	}

	for i := 0; i < len(inputList)-1; i++ {
		diff := math.Abs(float64(inputList[i] - inputList[i+1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func dropLevelCheck(inputList []int) bool {
	for i := 0; i < len(inputList); i++ {
		newList := []int{}
		for j, val := range inputList {
			if j != i {
				newList = append(newList, val)
			}
		}

		if checkCondition(newList) {
			return true
		}
	}

	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func openFile(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// two int lists
	var array [][]int

	// Read each line
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		parts := strings.Fields(line)
		row := make([]int, len(parts))
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			row[i] = num
		}

		array = append(array, row)

	}

	return array
}

func readExample() [][]int {
	return openFile("day2/example.txt")
}

func readInput() [][]int {
	return openFile("day2/input.txt")
}
