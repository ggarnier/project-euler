package main

import "fmt"

var cache = make(map[int]int)

func main() {
	fmt.Println(Sum(EvenNumbersFrom(FibonacciNumbersBelow(4000000))))
}

func Fibonacci(number int) int {
	value, ok := readCache(number)
	if ok {
		return value
	}

	if number <= 2 {
		updateCache(number, number)
		return number
	}

	value = Fibonacci(number-1) + Fibonacci(number-2)
	updateCache(number, value)
	return value
}

func readCache(number int) (int, bool) {
	cacheValue, ok := cache[number]
	return cacheValue, ok
}

func updateCache(number int, value int) {
	cache[number] = value
}

func FibonacciNumbersBelow(maxValue int) []int {
	maxNumber := 100
	result := make([]int, 0, maxNumber)

	for index := 0; index < maxNumber; index++ {
		fib := Fibonacci(index + 1)
		if fib > maxValue {
			break
		}

		result = result[:index+1]
		result[index] = fib
	}

	return result
}

func EvenNumbersFrom(list []int) []int {
	result := make([]int, 0, len(list))

	for index := 0; index < len(list); index++ {
		if list[index]%2 == 0 {
			result = result[:len(result)+1]
			result[len(result)-1] = list[index]
		}
	}

	return result
}

func Sum(list []int) int {
	sum := 0
	for index := 0; index < len(list); index++ {
		sum += list[index]
	}

	return sum
}
