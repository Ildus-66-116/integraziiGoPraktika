package main

import "fmt"

const size1 = 5
const newSize = 2

func maximum1(input [size1]int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func minimum1(input [size1]int) int {
	min := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

func wrongCalculation(input [size1]int) {
	for i := 0; i < 6; i++ {
		fmt.Println(input[i])
	}
}

func main() {

	numbers := [size1]int{10, 49, 48, 20, 39}

	fmt.Println(maximum1(numbers))

	fmt.Println(minimum1(numbers))

	wrongCalculation(numbers)

	newNumbers := [...]int{10, 11, 12, 13} // даем компилятору самому вычислить длину массива

	newNumberSmall := [...]int{10}

	var newSmaltest int
	var newSecondSmaltest int

	if len(newNumberSmall) >= 2 {
		newSmaltest := newNumberSmall[0]
		newSecondSmaltest := newNumberSmall[1]
	}
	fmt.Println(newSmaltest, newSecondSmaltest)
}
