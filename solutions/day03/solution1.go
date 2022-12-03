package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/juliangruber/go-intersect"
)

var lowers = "abcdefghijklmnopqrstuvwxyz"
var uppers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getScore(item string) int {

	iupper := strings.Index(uppers, item)
	if iupper > 0 {
		// found upper char
		return 27 + iupper
	}

	ilower := strings.Index(lowers, item)
	if ilower > 0 {
		// found lower char
		return 1 + ilower
	}
	panic(item)
}

func solve1(records []string) int {
	count := 0
	for i := 0; i < len(records); i++ {
		rs := records[i]
		rslen := len(rs)
		rslenhalf := rslen / 2
		for j := 0; j < rslenhalf; j++ {
			// cycle through first half and search in second half
			curitem := string(rs[j])
			p := strings.Index(rs[rslenhalf:], curitem)
			if p > -1 {
				count += getScore(curitem)
				break
			}
		}
	}
	return count
}

func solve2(records []string) int {
	count := 0
	for i := 0; i < len(records); i += 3 {
		rs1 := strings.Split(records[i], "")
		rs2 := strings.Split(records[i+1], "")
		rs3 := strings.Split(records[i+2], "")
		rt1 := intersect.Simple(rs1, rs2)
		rt2 := intersect.Simple(rt1, rs3)
		count += getScore(fmt.Sprint(rt2[0]))
	}
	return count
}

func main() {
	file, err := os.Open("inputs/day03/input1.txt")
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
