package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var debug string = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

type leyline struct {
	start  point
	end    point
	isLine bool
}

type point struct {
	x int
	y int
}

func makePont(input string) point {
	input = strings.ReplaceAll(input, " ", "")
	temp := strings.Split(input, ",")
	r, _ := strconv.Atoi(temp[0])
	c, _ := strconv.Atoi(temp[1])
	return point{r, c}
}

func parseInput(input string) []leyline {
	output := make([]leyline, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		current := strings.Split(scanner.Text(), "->")
		start := makePont(current[0])
		end := makePont(current[1])
		output = append(output, leyline{start: start,
			end:    end,
			isLine: start.x == end.x || start.y == end.y})
	}
	return output
}

func isOverLap(a leyline, b leyline) {
	// is there intersection in x level?
	// what about y level?

}

func findOverLapping(lines []leyline) int {
	counter := 0
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines)-1; j++ {
			if lines[i].isLine && lines[j].isLine {
				// how many point intersect between i and j

			}
		}
	}

	return counter
}

func main() {
	leylines := parseInput(debug)
	//fmt.Println(leylines)
	fmt.Println(findOverLapping(leylines))
}
