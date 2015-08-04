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

func (Suite) TestIsMultipleOf(c *check.C) {
	c.Assert(IsMultipleOf(10, 3), check.Equals, false)
	c.Assert(IsMultipleOf(9, 3), check.Equals, true)
}

func (Suite) TestListMultiplesBelow(c *check.C) {
	c.Assert(ListMultiplesBelow(10, 3), check.DeepEquals, []int{3, 6, 9})
}

func (Suite) TestSumMultiplesBelow(c *check.C) {
	c.Assert(SumMultiplesBelow(10, 3, 5), check.Equals, 23)
	c.Assert(SumMultiplesBelow(1000, 3, 5), check.Equals, 233168)
}
