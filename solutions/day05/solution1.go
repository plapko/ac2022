package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ship [9]string

func loadShip() {
	ship[0] = "QMGCL"
	ship[1] = "RDLCTFHG"
	ship[2] = "VJFNMTWR"
	ship[3] = "JFDVQP"
	ship[4] = "NFMSLBT"
	ship[5] = "RNVHCDP"
	ship[6] = "HCT"
	ship[7] = "GSJVZNHP"
	ship[8] = "ZFHG"
}

func moveShip(from int, to int) {
	// move one char
	stack := ship[from]
	lastchar := stack[len(stack)-1:]
	ship[from] = stack[:len(stack)-1]
	stacknew := ship[to]
	ship[to] = ship[to] + lastchar
	fmt.Printf("Moved %s from %s to %s, result: %s %s\n", lastchar, stack, stacknew, ship[from], ship[to])
}

func moveStack(from int, to int, count int) {
	// move multiple chars
	stack := ship[from]
	tail := stack[len(stack)-count:]
	ship[from] = stack[:len(stack)-count]
	stacknew := ship[to]
	ship[to] = ship[to] + tail
	fmt.Printf("Moved %s from %s to %s, result: %s %s\n", tail, stack, stacknew, ship[from], ship[to])
}

func printShip() {
	fmt.Printf("%#v\n", ship)
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return result
}

func solve1(records []string) int {
	count := 0
	for i := 0; i < len(records); i++ {
		// move 5 from 8 to 2
		if !strings.Contains(records[i], "move") {
			// skip the header, we did that one manually
			continue
		}
		bd := strings.Fields(records[i])
		move := toInt(bd[1])
		from := toInt(bd[3])
		to := toInt(bd[5])
		// fmt.Printf("Moving %d from %d to %d\n", move, from, to)
		for j := 1; j <= move; j++ {
			moveShip(from-1, to-1)
		}
		printShip()
	}
	return count
}

func solve2(records []string) int {
	count := 0
	for i := 0; i < len(records); i++ {
		// move 5 from 8 to 2
		if !strings.Contains(records[i], "move") {
			// skip the header, we did that one manually
			continue
		}
		bd := strings.Fields(records[i])
		move := toInt(bd[1])
		from := toInt(bd[3])
		to := toInt(bd[5])
		fmt.Printf("Moving %d from %d to %d\n", move, from, to)
		moveStack(from-1, to-1, move)
		printShip()
	}
	return count
}

func main() {
	file, err := os.Open("inputs/day05/input1.txt")
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

	loadShip()
	//fmt.Printf("Solve1: %d\n", solve1(records))
	printShip()
	fmt.Printf("Solve2: %d\n", solve2(records))
}
