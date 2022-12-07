package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var THRESHOLD = 100001
var DISK = 70000000
var UNUSED = 30000000
var folders = make(map[string]int)

func goDeeper(s string) []string {
	var result []string
	for k := range folders {
		if strings.Contains(k, s) && s != k && k != "/" {
			result = append(result, k)
		}
	}
	//fmt.Printf("%#v\n", result)
	return result
}

func goParents(s string) []string {
	var result []string
	for k := range folders {
		if strings.Contains(s, k) && s != k && k != "/" {
			result = append(result, k)
		}
	}
	//fmt.Printf("%#v\n", result)
	return result
}

func getMaxDepth() int {

	result := 0
	for k := range folders {
		c := strings.Count(k, "/")
		if c > result {
			result = c
			// fmt.Printf("max: %s (%d)\n", k, c)
		}
	}
	return result - 1
}

func getParent(s string) string {
	result := ""
	//fmt.Printf("GetParent %s=>", s)
	lastInd := 0
	lastInd = strings.LastIndex(s, "/")
	result = s[:lastInd]
	lastInd = strings.LastIndex(result, "/")
	result = result[:lastInd+1]
	// fmt.Printf("=> %s\n", result)
	return result
}

func getDepth(s string) int {
	result := strings.Count(s, "/") - 1
	// fmt.Printf("Get Depth: %s (%d)\n", s, result)
	return result
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return result
}

func solve1(records []string) int {

	cur := "/" // current folder, format /a/b/c
	for i := 0; i < len(records); i++ {
		rs := strings.Split(records[i], " ")

		// fmt.Printf("Current: '%s', logic: '%s'\n", cur, rs[0])
		if rs[0] == "$" {
			// executed command
			if rs[1] == "cd" {
				// change directory
				folder := rs[2]
				// fmt.Printf("..cd '%s'\n", folder)
				if folder == "/" {
					cur = folder
				} else if folder == ".." {
					// leave subfolder
					var lastInd int
					lastInd = strings.LastIndex(cur, "/")
					cur = cur[:lastInd]
					lastInd = strings.LastIndex(cur, "/")
					cur = cur[:lastInd+1]
				} else {
					// enter a subfolder
					cur += folder + "/"
				}

			}
			if rs[1] == "ls" {
				// list directory, irrelevant
			}
		} else {
			// output of previous command
			if rs[0] == "dir" {
				// show folders, irrelevant
			} else {
				// actual file sizes to scan
				size := toInt(rs[0])
				folders[cur] += size
			}
		}
	}
	// end load input

	fmt.Printf("Input: %#v\n", folders)
	fmt.Printf("Max depth: %d\n", getMaxDepth())

	// merge values from childrent to parents
	for i := getMaxDepth(); i > 0; i-- {
		// go through each depth level starting from i=8
		fmt.Printf("Starting Depth iteration %d/%d\n", i, getMaxDepth())
		for k, v := range folders {
			if getDepth(k) == i {
				folders[getParent(k)] += v
			}
		}
	}

	fmt.Printf("Merged: %#v\n", folders)

	// count final values
	count := 0
	for _, v := range folders {

		// ignore anything that's higher in value
		if v < THRESHOLD {
			count += v
		}
	}

	fmt.Printf("Solve1: %d\n", count)

	unused_space := DISK - folders["/"]
	delete_space := UNUSED - unused_space
	fmt.Printf("Unused space: %d / %d. Please delete %d\n", unused_space, UNUSED, delete_space)

	target := "/"
	for k, v := range folders {
		if v > delete_space {
			if v < folders[target] {
				target = k
			}
		}
	}
	fmt.Printf("Solve2: %s %d\n", target, folders[target])

	return count
}

func main() {
	file, err := os.Open("inputs/day07/input1.txt")
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
}
