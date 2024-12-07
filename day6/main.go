package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	example_map := readExample()
	input_map := readInput()

	example_map2 := readExample()
	input_map2 := readInput()

	answer := solution1(example_map)
	fmt.Println("Answer for example solution 1: ", answer)

	answer = solution1(input_map)
	fmt.Println("Answer for input solution 1: ", answer)

	answer = solution2(example_map2)
	fmt.Println("Answer for example solution 2: ", answer)

	answer = solution2(input_map2)
	fmt.Println("Answer for input solution 2: ", answer)

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

	copyMap := make([][]string, len(data)) // Outer slice
	for i := range data {
		copyMap[i] = make([]string, len(data[i])) // Inner slice for each row
		copy(copyMap[i], data[i])                 // Copy each row
	}

	copyMap = walkMap(copyMap, x, y)

	// printMap(data)

	// first get all locations of current path minus starting location
	walkedPositions := getWalkedPositions(copyMap, POSTIONCHAR, x, y)

	// fmt.Println(walkedPositions)

	total := 0

	var wg sync.WaitGroup // WaitGroup to synchronize goroutines
	var mu sync.Mutex     // Mutex to protect the shared variable 'total'

	// Launch goroutines
	for i := 0; i < len(walkedPositions); i++ {
		wg.Add(1) // Increment the WaitGroup counter

		go func(i int) {
			defer wg.Done() // Decrement the WaitGroup counter when done

			// Create a copy of the 2D slice
			copyMap := make([][]string, len(data))
			for j := range data {
				copyMap[j] = make([]string, len(data[j]))
				copy(copyMap[j], data[j])
			}

			// Call determineCircle
			result := determineCircle(copyMap, walkedPositions[i], x, y, i)

			// Safely update the shared 'total' variable
			mu.Lock()
			total += result
			mu.Unlock()
		}(i) // Pass 'i' to avoid closure capturing issues
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return total
}

func determineCircle(data [][]string, position []int, x, y, i int) int {
	data[position[0]][position[1]] = "#"
	walkedRecord := make([][]int, 0)
	BLOCKCHAR := "#"
	secondRound := false
	nextStep := []int{0, 0}

	direction := "up"
	count := 0

	// postion to string

	//convert []int to string in one line

	// file, _ := os.Create(strconv.Itoa(i) + "_" + strconv.Itoa(position[0]) + "_" + strconv.Itoa(position[1]) + ".txt")
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
		if count == 10000 {
			return 1
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
			switch direction {
			case "up":
				data[x][y] = "|"
			case "down":
				data[x][y] = "|"
			case "left":
				data[x][y] = "-"
			case "right":
				data[x][y] = "-"
			}

			if secondRound {
				if nextStep[0] == x && nextStep[1] == y {
					// file.Close()
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
		// file.WriteString("count:" + strconv.Itoa(count) + "\n")
		// for i := 0; i < len(data); i++ {
		// 	for j := 0; j < len(data[i]); j++ {
		// 		file.WriteString(data[i][j])
		// 	}
		// 	file.WriteString("\n")
		// }
		// file.WriteString("\n")

		count++
	}
	// file.Close()

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
