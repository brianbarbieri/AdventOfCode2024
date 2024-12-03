package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func solution1(data string) int {
	searchState, firstNumber, secondNumber := "m", "", ""
	numbers := [][]int{}

	for _, c := range data {

		switch c {
		case 'm':
			if searchState == "m" {
				searchState = "u"
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case 'u':
			if searchState == "u" {
				searchState = "l"
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case 'l':
			if searchState == "l" {
				searchState = "("
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case '(':
			if searchState == "(" {
				searchState = "firstNum"
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case ',':
			if searchState == "firstNum" {
				searchState = "secondNum"
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case ')':
			if searchState == "secondNum" {
				searchState = "m"
				numbers = append(numbers, []int{toInt(firstNumber), toInt(secondNumber)})
				searchState, firstNumber, secondNumber = "m", "", ""
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		default:
			if unicode.IsDigit(c) {
				if searchState == "firstNum" {
					firstNumber += string(c)
				} else if searchState == "secondNum" {
					secondNumber += string(c)
				}
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		}
	}
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i][0] * numbers[i][1]
	}
	return total
}

func solution2(data string) int {
	searchState, firstNumber, secondNumber, doState := "m", "", "", true
	numbers := [][]int{}
	data = strings.Replace(data, "do()", "=", -1)    // no occurences of "=" in the data so replace it to make parsing easier
	data = strings.Replace(data, "don't()", ".", -1) // same for .
	for _, c := range data {
		switch c {
		case '=':
			doState = true
		case '.':
			doState = false
		case 'm':
			if searchState == "m" {
				searchState = "u"
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case 'u':
			if searchState == "u" {
				searchState = "l"
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case 'l':
			if searchState == "l" {
				searchState = "("
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case '(':
			if searchState == "(" {
				searchState = "firstNum"
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case ',':
			if searchState == "firstNum" {
				searchState = "secondNum"
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		case ')':
			if searchState == "secondNum" {
				searchState = "m"
				if doState {
					numbers = append(numbers, []int{toInt(firstNumber), toInt(secondNumber)})
				}
				searchState, firstNumber, secondNumber = "m", "", ""
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		default:
			if unicode.IsDigit(c) {
				if searchState == "firstNum" {
					firstNumber += string(c)
				} else if searchState == "secondNum" {
					secondNumber += string(c)
				}
			} else {
				searchState, firstNumber, secondNumber = "m", "", ""
			}
		}
	}
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i][0] * numbers[i][1]
	}
	return total
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func openFile(filePath string) string {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		return line

	}

	return ""
}

func readExample() string {
	return openFile("day3/example.txt")
}

func readExample2() string {
	return openFile("day3/example2.txt")
}

func readInput() string {
	return openFile("day3/input.txt")
}
