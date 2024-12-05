package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data_example_rules, data_example_update := readExample()
	data_input_rules, data_input_update := readInput()

	answer := solution1(data_example_rules, data_example_update)
	fmt.Println("Answer for example solution 1: ", answer)

	answer = solution1(data_input_rules, data_input_update)
	fmt.Println("Answer for input solution 1: ", answer)

	answer = solution2(data_example_rules, data_example_update)
	fmt.Println("Answer for example solution 2: ", answer)

	answer = solution2(data_input_rules, data_input_update)
	fmt.Println("Answer for input solution 2: ", answer)

}

type rule struct {
	numsBefore, numsAfter []int
}

func solution2(rules [][]int, update [][]int) int {
	ruleMap := makeRuleMap(rules)
	total := 0

	var incorrelyOrdered [][]int
	for _, updateLine := range update {
		if isCorrectOrder(ruleMap, updateLine) == 0 {
			incorrelyOrdered = append(incorrelyOrdered, updateLine)
		}
	}

	for _, line := range incorrelyOrdered {
		line = orderLine(ruleMap, line)
		total += line[len(line)/2]
	}
	return total
}

func solution1(rules [][]int, update [][]int) int {
	ruleMap := makeRuleMap(rules)
	total := 0

	for _, updateLine := range update {
		total += isCorrectOrder(ruleMap, updateLine)
	}

	return total
}

func orderLine(orderMap map[int]rule, line []int) []int {
	for i, num := range line {
		afterNums := orderMap[num].numsBefore
		beforeNums := orderMap[num].numsAfter
		if len(afterNums) == 0 && len(beforeNums) == 0 {
			continue
		}
		before := line[:i]
		after := line[i+1:]
		var wrongSpot []int
		for _, bn := range beforeNums {
			for _, a := range after {
				if bn == a {
					wrongSpot = append(wrongSpot, bn)
				}
			}
		}

		if len(wrongSpot) > 0 {
			var newAfter []int
			for _, a := range after {
				found := false
				for _, ws := range wrongSpot {
					if ws == a {
						found = true
						break
					}
				}
				if !found {
					newAfter = append(newAfter, a)
				}
			}
			line = append(append(append(before, wrongSpot...), num), newAfter...)
			return orderLine(orderMap, line)
		}
	}
	return line
}

func isCorrectOrder(ruleMap map[int]rule, updateLine []int) int {
	for i, num := range updateLine {
		rule := ruleMap[num]

		before := updateLine[:i]
		for _, beforeNum := range before {
			if contains(rule.numsBefore, beforeNum) {
				return 0
			}
		}

		after := updateLine[i+1:]
		for _, afterNum := range after {
			if contains(rule.numsAfter, afterNum) {
				return 0
			}
		}

	}
	return updateLine[len(updateLine)/2]

}

func makeRuleMap(rules [][]int) map[int]rule {
	ruleMap := make(map[int]rule)

	for _, r := range rules {
		rule0 := ruleMap[r[0]]
		rule0.numsBefore = append(rule0.numsBefore, r[1])
		ruleMap[r[0]] = rule0
		rule1 := ruleMap[r[1]]
		rule1.numsAfter = append(rule1.numsAfter, r[0])
		ruleMap[r[1]] = rule1
	}

	return ruleMap
}

func contains(list []int, beforeNum int) bool {
	for _, num := range list {
		if num == beforeNum {
			return true
		}
	}
	return false
}

func openFile(filePath string) ([][]int, [][]int) {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules [][]int
	var update [][]int
	reachedUpdates := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			reachedUpdates = true
			continue
		}
		if !reachedUpdates {
			parts := strings.Split(line, "|")
			rules = append(rules, []int{toInt(parts[0]), toInt(parts[1])})
		} else {
			parts := strings.Split(line, ",")
			var updateLine []int
			for _, part := range parts {
				updateLine = append(updateLine, toInt(part))
			}
			update = append(update, updateLine)
		}
	}

	return rules, update
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func readExample() ([][]int, [][]int) {
	return openFile("day5/example.txt")
}

func readInput() ([][]int, [][]int) {
	return openFile("day5/input.txt")
}
