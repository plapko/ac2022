package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/juliangruber/go-intersect"
)

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return result
}

func solve1(records []string) int {
	count1 := 0
	count2 := 0
	for i := 0; i < len(records); i++ {
		sections := strings.Split(records[i], ",")
		sec1 := strings.Split(sections[0], "-") // 2, 3
		sec2 := strings.Split(sections[1], "-") // 4, 5
		str1 := ""
		str2 := ""
		for j := 1; j <= toInt(sec1[1]); j++ {
			// cycle through first pair and make a string
			if j < toInt(sec1[0]) {
				str1 += ""
			} else {
				str1 += "("
				str1 += fmt.Sprint(j)
				str1 += ")"
			}
		}
		for j := 1; j <= toInt(sec2[1]); j++ {
			// cycle through second pair and make a string
			if j < toInt(sec2[0]) {
				str2 += ""
			} else {
				str2 += "("
				str2 += fmt.Sprint(j)
				str2 += ")"
			}
		}
		if strings.Contains(str1, str2) {
			count1++
			continue
		}
		if strings.Contains(str2, str1) {
			count2++
			continue
		}
	}
	return count1 + count2
}

func solve2(records []string) int {
	count1 := 0
	count2 := 0

	for i := 0; i < len(records); i++ {
		sections := strings.Split(records[i], ",")
		sec1 := strings.Split(sections[0], "-") // 2, 3
		sec2 := strings.Split(sections[1], "-") // 4, 5
		var sl1 []int
		var sl2 []int
		for j := 1; j <= toInt(sec1[1]); j++ {
			// cycle through first pair and make a slice
			if j >= toInt(sec1[0]) {
				sl1 = append(sl1, j)
			}
		}
		for j := 1; j <= toInt(sec2[1]); j++ {
			// cycle through second pair and make a slice
			if j >= toInt(sec2[0]) {
				sl2 = append(sl2, j)
			}
		}
		if len(intersect.Simple(sl1, sl2)) > 0 {
			count1++
			continue
		}
		if len(intersect.Simple(sl2, sl1)) > 0 {
			count2++
			continue
		}
	}

	return count1 + count2
}

func main() {
	file, err := os.Open("inputs/day04/input1.txt")
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
