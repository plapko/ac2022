package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve1(records []string) int {
	count := 0
	signal := records[0]
	for i := 4; i < len(signal); i++ {
		last4 := signal[i-4 : i]
		good := true
		fmt.Printf("%d: %s", i, last4)
		for j := 0; j <= 3; j++ {
			fmt.Printf(" %s", last4[j:j+1])

			if strings.Count(last4, last4[j:j+1]) > 1 {
				good = false
			}
		}
		fmt.Println()
		if good {
			count = i
			fmt.Printf("Good index: %d\n", i)
			break
		}
	}
	return count
}

func solve2(records []string, marker int) int {
	count := 0
	signal := records[0]
	for i := marker; i < len(signal); i++ {
		last4 := signal[i-marker : i]
		good := true
		fmt.Printf("%d: %s", i, last4)
		for j := 0; j <= marker-1; j++ {
			fmt.Printf(" %s", last4[j:j+1])

			if strings.Count(last4, last4[j:j+1]) > 1 {
				good = false
			}
		}
		fmt.Println()
		if good {
			count = i
			fmt.Printf("Good index: %d\n", i)
			break
		}
	}
	return count
}

func main() {
	file, err := os.Open("inputs/day06/input1.txt")
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
	fmt.Printf("Solve2: %d\n", solve2(records, 14))
}
