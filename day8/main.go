package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Point struct {
	X, Y int
}

// Parse the map and find antenna positions and frequencies
func parseMap(grid []string) map[byte][]Point {
	antennas := make(map[byte][]Point)
	for y, line := range grid {
		for x, char := range line {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				antennas[byte(char)] = append(antennas[byte(char)], Point{X: x, Y: y})
			}
		}
	}
	return antennas
}

// Calculate unique antinodes based on the original rules
func calculateAntinodes(antennas map[byte][]Point, rows, cols int) map[Point]struct{} {
	antinodes := make(map[Point]struct{})

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1, p2 := points[i], points[j]

				// Calculate midpoint
				mid := Point{X: (p1.X + p2.X) / 2, Y: (p1.Y + p2.Y) / 2}

				// Check if the midpoint is a valid antinode
				if (p1.X+p2.X)%2 == 0 && (p1.Y+p2.Y)%2 == 0 {
					antinodes[mid] = struct{}{}
				}

				// Calculate extrapolated antinodes
				dx := p2.X - p1.X
				dy := p2.Y - p1.Y

				// Extrapolated points
				extrapolated1 := Point{X: p1.X - dx, Y: p1.Y - dy}
				extrapolated2 := Point{X: p2.X + dx, Y: p2.Y + dy}

				// Add extrapolated points if within bounds
				if extrapolated1.X >= 0 && extrapolated1.X < cols && extrapolated1.Y >= 0 && extrapolated1.Y < rows {
					antinodes[extrapolated1] = struct{}{}
				}
				if extrapolated2.X >= 0 && extrapolated2.X < cols && extrapolated2.Y >= 0 && extrapolated2.Y < rows {
					antinodes[extrapolated2] = struct{}{}
				}
			}
		}
	}

	return antinodes
}

// Calculate unique antinodes based on the updated rules (solution2)
func solution2(antennas map[byte][]Point, rows, cols int) map[Point]struct{} {
	antinodes := make(map[Point]struct{})

	// Iterate over each frequency and its corresponding antenna positions
	for _, points := range antennas {
		// For each antenna, mark its position as an antinode
		for _, p1 := range points {
			antinodes[p1] = struct{}{} // Mark the antenna position itself as an antinode

			// Check all other antennas of the same frequency
			for _, p2 := range points {
				if p1 == p2 {
					continue
				}

				// Check if p1 and p2 are aligned (same row or column)
				if p1.X == p2.X || p1.Y == p2.Y {
					// If aligned in a column (same X), add all points between p1 and p2 in that column
					if p1.X == p2.X {
						minY, maxY := p1.Y, p2.Y
						if p1.Y > p2.Y {
							minY, maxY = p2.Y, p1.Y
						}
						for y := minY; y <= maxY; y++ {
							antinodes[Point{X: p1.X, Y: y}] = struct{}{}
						}
					}
					// If aligned in a row (same Y), add all points between p1 and p2 in that row
					if p1.Y == p2.Y {
						minX, maxX := p1.X, p2.X
						if p1.X > p2.X {
							minX, maxX = p2.X, p1.X
						}
						for x := minX; x <= maxX; x++ {
							antinodes[Point{X: x, Y: p1.Y}] = struct{}{}
						}
					}
				}
			}
		}
	}

	return antinodes
}

func main() {
	grid := readExample()
	grid2 := readExample()
	gridi := readInput()
	gridi2 := readInput()

	rows, cols := len(grid), len(grid[0])
	rows2, cols2 := len(gridi), len(gridi[0])

	antennas1 := parseMap(grid)
	antennas2 := parseMap(gridi)
	antennas3 := parseMap(grid2)
	antennas4 := parseMap(gridi2)

	antinodes1 := calculateAntinodes(antennas1, rows, cols)
	fmt.Println("Solution 1: Total unique antinode locations:", len(antinodes1))

	antinodes1 = calculateAntinodes(antennas2, rows2, cols2)
	fmt.Println("Solution 2: Total unique antinode locations:", len(antinodes1))

	antinodes2 := solution2(antennas3, rows, cols)
	fmt.Println("Solution 1: Total unique antinode locations:", len(antinodes2))

	antinodes2 = solution2(antennas4, rows2, cols2)
	fmt.Println("Solution 2: Total unique antinode locations:", len(antinodes2))
}

func openFile(filePath string) []string {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, line)
	}

	return grid
}

func readExample() []string {
	return openFile("day8/example.txt")
}

func readInput() []string {
	return openFile("day8/input.txt")
}
