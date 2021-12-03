package main

import (
	"fmt"
	"math"
	"misc"
)

func binaryToDecimal(num []int) int {
	output := 0
	for i := 0; i < len(num); i++ {
		b := num[i]
		output += b * int(math.Pow(2, float64(len(num)-i-1)))
	}
	return output
}

func getPowerConsumption(arr []string) int {

	temp := make([]int, len(arr[0]))
	gamma := make([]int, len(arr[0]))
	epsilon := make([]int, len(arr[0]))

	for i := range temp {
		temp[i] = 0
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == '1' {
				temp[j]++
			}
		}
	}

	for i := range temp {
		if temp[i] > len(arr)/2 {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}
	return binaryToDecimal(gamma) * binaryToDecimal(epsilon)
}

func filter(arr []string, pos int, reverse bool) []string {
	counter := 0
	majorityOne := true
	output := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i][pos] == '1' {
			counter++
		}
	}

	if !reverse {
		majorityOne = counter >= len(arr)-counter
	} else {
		majorityOne = counter < len(arr)-counter
	}

	for i := 0; i < len(arr); i++ {
		if majorityOne && arr[i][pos] == '1' {
			output = append(output, arr[i])
		} else if !majorityOne && arr[i][pos] == '0' {
			output = append(output, arr[i])
		}
	}

	return output
}

func getLifeSupportRating(arr []string) int {
	arr1 := arr
	arr2 := arr

	for i := 0; i < len(arr[0]); i++ {
		if len(arr1) != 1 {
			arr1 = filter(arr1, i, false)
		}
		if len(arr2) != 1 {
			arr2 = filter(arr2, i, true)
		}
		if len(arr1) == 1 && len(arr2) == 1 {
			break
		}
	}

	oxygen := make([]int, len(arr[0]))
	scubber := make([]int, len(arr[0]))
	for i := 0; i < len(arr[0]); i++ {
		oxygen[i] = int(arr1[0][i] - '0')
		scubber[i] = int(arr2[0][i] - '0')
	}

	return binaryToDecimal(oxygen) * binaryToDecimal(scubber)
}

func main() {
	arr := misc.GetInput(3)
	fmt.Println(getPowerConsumption(arr))
	fmt.Println(getLifeSupportRating(arr))
}
