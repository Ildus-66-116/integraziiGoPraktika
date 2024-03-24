package main

import "fmt"

// 10 49 48 20 39 Massiv
const size = 5

func maximum(input [size]int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func minimum(input [size]int) int {
	min := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

func main() {
	//var numbers [5]int
	//			  0   1   2   3   4
	numbers := [size]int{10, 49, 48, 20, 39}

	//fmt.Println(numbers[0], numbers[1])

	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}
	fmt.Println(maximum(numbers))
	fmt.Println(minimum(numbers))
}
