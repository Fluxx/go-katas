package main

import (
	"fmt"
	"math"
)

const Precision = 1e-15

func Newton(guess, input float64) float64 {
	return guess - (guess*guess-input)/(2*guess)
}

func Sqrt(input float64) float64 {
	var guess, delta float64 = input, 1

	for delta > Precision {
		// fmt.Println(guess, delta)
		old_guess := guess

		guess = Newton(guess, input)
		delta = math.Abs(guess - old_guess)
	}

	return guess
}

func main() {
	fmt.Println(9, Sqrt(9))
	fmt.Println(25, Sqrt(25))
	fmt.Println(100, Sqrt(100))
	fmt.Println(52, Sqrt(52))
}
