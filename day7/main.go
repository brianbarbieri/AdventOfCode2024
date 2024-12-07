package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	answer  int
	numbers []int
}

func main() {
	data_example := readExample()
	data_input := readInput()

	answer := solution1(data_example)
	fmt.Println("Answer for example solution 1: ", answer)

	answer = solution1(data_input)
	fmt.Println("Answer for input solution 1: ", answer)

	answer = solution2(data_example)
	fmt.Println("Answer for example solution 2: ", answer)

	answer = solution2(data_input)
	fmt.Println("Answer for input solution 2: ", answer)
}

func solution1(data []Line) int {
	total := 0
	for _, line := range data {
		answers := []int{line.numbers[0]}
		for _, num := range line.numbers[1:] {
			newAnswers := []int{}
			for _, answer := range answers {
				newAnswers = append(newAnswers, answer+num)
				newAnswers = append(newAnswers, answer*num)
			}
			answers = newAnswers
		}

		for _, answer := range answers {
			if answer == line.answer {
				total += answer
				break
			}
		}
	}
	return total
}

func solution2(data []Line) int {
	total := 0
	for _, line := range data {
		answers := []int{line.numbers[0]}
		for _, num := range line.numbers[1:] {
			newAnswers := []int{}
			for _, answer := range answers {
				newAnswers = append(newAnswers, answer+num)
				newAnswers = append(newAnswers, answer*num)
				strAns := strconv.Itoa(answer) + strconv.Itoa(num)
				num, _ := strconv.Atoi(strAns)
				newAnswers = append(newAnswers, num)
			}
			answers = newAnswers
		}

		for _, answer := range answers {
			if answer == line.answer {
				total += answer
				break
			}
		}
	}
	return total
}

func openFile(filePath string) []Line {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var array []Line

	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		parts := strings.Split(line, ": ")
		ans, numbers := parts[0], parts[1]

		parts = strings.Fields(numbers)
		row := make([]int, len(parts))
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			row[i] = num
		}
		answer, _ := strconv.Atoi(ans)
		array = append(array, Line{answer: answer, numbers: row})

	}

	return array
}

func readExample() []Line {
	return openFile("day7/example.txt")
}

func readInput() []Line {
	return openFile("day7/input.txt")
}
