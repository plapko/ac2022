package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// A = Rock		X = Rock		Loss
// B = Paper	Y = Paper		Draw
// C = Scisors	Z = Scissors	Win

func solve2(records []string) int {
	score := 0
	for i := 0; i < len(records); i++ {
		words := strings.Fields(records[i])
		g1 := words[0] // ABC
		g2 := words[1] // XYZ

		// Loss
		if g2 == "X" {
			score += 0 // Loss
			if g1 == "A" {
				score += 3 // Z
			}
			if g1 == "B" {
				score += 1 // X
			}
			if g1 == "C" {
				score += 2 // Y
			}
		}

		// Draw
		if g2 == "Y" {
			score += 3 // Draw
			if g1 == "A" {
				score += 1 // X
			}
			if g1 == "B" {
				score += 2 // Y
			}
			if g1 == "C" {
				score += 3 // Z
			}
		}

		// Win
		if g2 == "Z" {
			score += 6 // Win
			if g1 == "A" {
				score += 2 // Y
			}
			if g1 == "B" {
				score += 3 // Z
			}
			if g1 == "C" {
				score += 1 // X
			}
		}

		// fmt.Printf("%s %s\n", g1, g2)
	}
	return score
}

func solve1(records []string) int {
	score := 0
	for i := 0; i < len(records); i++ {
		words := strings.Fields(records[i])
		g1 := words[0] // ABC
		g2 := words[1] // XYZ

		// draw
		if g1 == "A" && g2 == "X" {
			score += 3 + 1
		}
		if g1 == "B" && g2 == "Y" {
			score += 3 + 2
		}
		if g1 == "C" && g2 == "Z" {
			score += 3 + 3
		}

		// opponent wins
		if g1 == "A" && g2 == "Z" {
			score += 3 + 0
		}
		if g1 == "B" && g2 == "X" {
			score += 1 + 0
		}
		if g1 == "C" && g2 == "Y" {
			score += 2 + 0
		}

		// I win
		if g1 == "A" && g2 == "Y" {
			score += 2 + 6
		}
		if g1 == "B" && g2 == "Z" {
			score += 3 + 6
		}
		if g1 == "C" && g2 == "X" {
			score += 1 + 6
		}

		// fmt.Printf("%s %s\n", g1, g2)
	}
	return score
}

func main() {
	file, err := os.Open("inputs/day02/input1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var records []string

	for scanner.Scan() {
		records = append(records, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Solve1: %d\n", solve1(records))
	fmt.Printf("Solve2: %d\n", solve2(records))
}
