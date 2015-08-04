package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SumMultiplesBelow(1000, 3, 5))
}

func IsMultipleOf(number int, base int) bool {
	return math.Mod(float64(number), float64(base)) == 0
}

func ListMultiplesBelow(topValue int, base int) []int {
	result := make([]int, 0)

	for i := 1; i < topValue; i++ {
		if IsMultipleOf(i, base) {
			result = append(result, i)
		}
	}

	return result
}

func SumMultiplesBelow(topValue int, base1 int, base2 int) int {
	list1 := ListMultiplesBelow(topValue, base1)
	list2 := ListMultiplesBelow(topValue, base2)

	fullList := make([]int, len(list1))
	copy(fullList, list1)

	for _, number := range list2 {
		if !intInSlice(number, list1) {
			fullList = append(fullList, number)
		}
	}

	return sumNumbers(fullList)
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func sumNumbers(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}
