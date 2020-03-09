package lab1

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrefixToInfix(c *C) {
	var res string
	var err error
	// Difficult exampl
	res, _ = PrefixToInfix("+-*/ABCD/E/F+GH")
	c.Assert(res, Equals, "((((A/B)*C)-D)+(E/(F/(G+H))))")
	// Simple examples
	res, _ = PrefixToInfix("-+ABC")
	c.Assert(res, Equals, "((A+B)-C)")
	// Error case "must be non empty string"
	_, err = PrefixToInfix("")
	c.Assert(err, ErrorMatches, "Input must be non empty string")
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("*+AB-CD")
	fmt.Println(res)
}
