package main

import (
	"fmt"
	"math"
)

func ConcatenateStrings(a, b string) string {
	return a + b
}

func AddIntegers(a, b int) int {
	return a + b
}

func DivideFloats(a, b float64) float64 {
	return math.Floor((a/b)*100) / 100
}

func And(a, b bool) bool {
	return a && b
}

func Or(a, b bool) bool {
	return a || b
}

func Not(a bool) bool {
	return !a
}

func main() {
	fmt.Println(ConcatenateStrings("Go", "lang"))

	fmt.Println(1, "+", 1, " = ", AddIntegers(1, 1))
	fmt.Println(7, "/", 3, " = ", DivideFloats(7.0, 3.0))

	fmt.Println(And(true, false))
	fmt.Println(Or(true, false))
	fmt.Println(Not(true))
}
