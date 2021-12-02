package main

import (
	"fmt"
	"misc"
	"strconv"
)

func getAmountOfIncreases(arr []string) int {
	var prev *int = nil
	var counter int = 0

	for i := 0; i < len(arr); i++ {
		cur, _ := strconv.Atoi(arr[i])
		if prev != nil && cur > *prev {
			counter++
		}
		prev = &cur

	}
	return counter
}

func getIncreaseInThree(arr []string) int {
	var prev *int = nil
	var counter int = 0

	for i := 0; i < len(arr)-2; i++ {
		var cur int = 0
		for j := 0; j < 3; j++ {
			temp, _ := strconv.Atoi(arr[i+j])
			cur += temp
		}
		if prev != nil && cur > *prev {
			counter++
		}
		prev = &cur
	}
	return counter
}

func main() {

	arr := misc.GetInput(1)
	fmt.Println(arr)
	fmt.Println(getAmountOfIncreases(arr))
	fmt.Println(getIncreaseInThree(arr))
}
