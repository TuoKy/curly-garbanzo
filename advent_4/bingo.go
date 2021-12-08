package main

import (
	"bufio"
	"fmt"
	"misc"
	"strconv"
	"strings"
)

var lenght int = 5

type board struct {
	rows   []bingoNumber
	points int
}

type bingoNumber struct {
	value    int
	selected bool
}

func (r *board) RoundsToWin(nums []int) int {
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if r.Select(n) && r.Bingo() {
			sum := 0
			for _, k := range r.rows {
				if !k.selected {
					sum += k.value
				}
			}
			r.points = sum * n
			return i
		}
	}
	return 0
}

func (r *board) Bingo() bool {
	for i := 0; i < lenght; i++ {
		counterY := 0
		counterX := 0

		// rows
		for j := 0; j < lenght; j++ {
			if r.rows[i+j*lenght].selected {
				counterY++
			} else {
				break
			}
		}

		// columns
		for j := 0; j < lenght; j++ {
			if r.rows[i*lenght+j].selected {
				counterX++
			} else {
				break
			}
		}
		if counterX == 5 || counterY == 5 {
			return true
		}
	}
	return false
}

func (r *board) Select(i int) bool {
	for j := 0; j < len(r.rows); j++ {
		if r.rows[j].value == i {
			r.rows[j].selected = true
			return true
		}
	}
	return false
}

func convert(param string) bingoNumber {
	return bingoNumber{func(n int, _ error) int { return n }(strconv.Atoi(param)), false}
}

func parseInput(input string) ([]int, []board) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	first := true
	boardIndex := 0

	numbers := make([]int, 0)
	boards := make([]board, 0)
	boards = append(boards, board{make([]bingoNumber, 0), 0})
	for scanner.Scan() {
		current := strings.Fields(scanner.Text())
		if first {
			// First line should be the random numbers
			current = strings.Split(current[0], ",")
			for _, c := range current {
				numbers = append(numbers, func(n int, _ error) int { return n }(strconv.Atoi(c)))
			}
			first = false
			scanner.Scan() // skip next line change
		} else {
			// Rest are bingo boards seperated by empty line
			if len(current) < 1 {
				boardIndex++
				boards = append(boards, board{make([]bingoNumber, 0), 0})
			} else {
				for _, c := range current {
					boards[boardIndex].rows = append(boards[boardIndex].rows, convert(c))
				}
			}
		}

	}
	return numbers, boards
}

func calculateRounds(numbers []int, boards []board) []int {
	output := make([]int, len(boards))
	for i := 0; i < len(boards); i++ {
		output[i] = boards[i].RoundsToWin(numbers)
	}
	return output
}

func getFirstWinningBoardScore(rounds []int) int {
	var cur *int = nil
	output := 0
	for i := 0; i < len(rounds); i++ {
		r := rounds[i]
		if cur == nil || *cur > r {
			cur = &r
			output = i
		}
	}
	return output
}

func getLastWinningBoardScore(rounds []int) int {
	var cur *int = nil
	output := 0
	for i := 0; i < len(rounds); i++ {
		r := rounds[i]
		if cur == nil || *cur < r {
			cur = &r
			output = i
		}
	}
	return output
}

func main() {
	arr := misc.GetRawInput(4)
	numbers, boards := parseInput(arr)
	rounds := calculateRounds(numbers, boards)
	fmt.Println("Part 1:", boards[getFirstWinningBoardScore(rounds)].points)
	fmt.Println("Part 2:", boards[getLastWinningBoardScore(rounds)].points)
}

var debug string = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`
