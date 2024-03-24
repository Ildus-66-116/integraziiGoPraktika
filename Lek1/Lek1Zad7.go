package main

import (
	"fmt"
	"math"
)

const size2 = 10

func twoSmaltestNumbers(input [size2]int) (int, int) {
	min := input[0]

	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}

	secondMin := math.MaxInt
	for i := 0; i < len(input); i++ {
		if input[i] < secondMin && input[i] > min {
			secondMin = input[i]
		}
	}
	return min, secondMin
}

func newTwoSmaltestNumbers(input [size2]int) (int, int) {
	min := math.MaxInt
	secondMin := math.MaxInt
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			secondMin = min
			min = input[i]
		} else if input[i] < secondMin {
			secondMin = input[i]
		}
	}
	return min, secondMin
}

func newTwoLangerstestNumbers(input [size2]int) (int, int) {
	max := math.MinInt
	secondMax := math.MinInt
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			secondMax = max
			max = input[i]
		} else if input[i] > secondMax {
			secondMax = input[i]
		}
	}
	return max, secondMax
}

func main() {
	numbers := [size2]int{10, 20, 30, 40, 50, 90, 80, 70, 60, 2}
	fmt.Println(twoSmaltestNumbers(numbers))
	fmt.Println(newTwoSmaltestNumbers(numbers))
	fmt.Println(newTwoLangerstestNumbers(numbers))
}
