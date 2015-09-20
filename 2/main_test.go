package main

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

var _ = check.Suite(Suite{})

type Suite struct{}

func (Suite) TestFibonacci(c *check.C) {
	c.Assert(Fibonacci(1), check.Equals, 1)
	c.Assert(Fibonacci(2), check.Equals, 2)
	c.Assert(Fibonacci(3), check.Equals, 3)
	c.Assert(Fibonacci(8), check.Equals, 34)
}

func (Suite) TestFibonacciNumbersBelow4Million(c *check.C) {
	result := FibonacciNumbersBelow(4000000)
	c.Assert(len(result), check.Equals, 32)
	c.Assert(result[31], check.Equals, 3524578)
}

func (Suite) TestEvenNumbersFrom(c *check.C) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := EvenNumbersFrom(numbers)
	c.Assert(len(result), check.Equals, 5)
	c.Assert(result, check.DeepEquals, []int{2, 4, 6, 8, 10})
}

func (Suite) TestSum(c *check.C) {
	c.Assert(Sum([]int{1, 2, 3, 4}), check.Equals, 10)
	c.Assert(Sum([]int{7, 10, 9}), check.Equals, 26)
}
