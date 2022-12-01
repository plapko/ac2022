package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
		// panic(err)
	}
	return result
}

func solve2(records []int) int {
	summations := make(map[int]int)
	j := 0
	for i := 0; i < len(records); i++ {
		summations[j] = records[i] + summations[j]
		if records[i] == 0 {
			j++
		}
	}
	type kv struct {
		Key   int
		Value int
	}
	var ss []kv
	for k, v := range summations {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	return ss[0].Value + ss[1].Value + ss[2].Value
}

func solve1(records []int) int {
	summations := make(map[int]int)
	j := 0
	max := 0
	for i := 0; i < len(records); i++ {
		summations[j] = records[i] + summations[j]
		if summations[j] > max {
			max = summations[j]
		}
		if records[i] == 0 {
			j++
		}
	}
	return max
}

func main() {
	file, err := os.Open("inputs/day01/input1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var records []int

	for scanner.Scan() {
		records = append(records, toInt(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	//fmt.Printf("Solve1: %d\n", solve1(records))
	fmt.Printf("Solve2: %d\n", solve2(records))
}
