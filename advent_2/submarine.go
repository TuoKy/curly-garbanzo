package main

import (
	"fmt"
	"misc"
	"strconv"
)

func getLocationProduct(arr []string) int {
	var y int = 0
	var x int = 0

	for i := 0; i < len(arr); i += 2 {
		v, _ := strconv.Atoi(arr[i+1])
		switch arr[i] {
		case "up":
			y -= v
		case "down":
			y += v
		case "forward":
			x += v
		}
	}

	return x * y
}

func getLocationProductWithAim(arr []string) int {
	var y int = 0
	var x int = 0
	var aim int = 0

	for i := 0; i < len(arr); i += 2 {
		v, _ := strconv.Atoi(arr[i+1])
		switch arr[i] {
		case "up":
			aim -= v
		case "down":
			aim += v
		case "forward":
			x += v
			y += v * aim
		}
	}

	return x * y
}

func main() {

	arr := misc.GetInput(2)

	//fmt.Println(arr)
	fmt.Println(getLocationProduct(arr))
	fmt.Println(getLocationProductWithAim(arr))
}
