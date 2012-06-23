package main

import (
	"fmt"
	"math"
)

const Precision = 1e-15

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Newton(guess, input float64) float64 {
	return guess - (guess*guess-input)/(2*guess)
}

func Sqrt(input float64) (float64, error) {
	var guess, delta float64 = input, 1

	if guess < 0 {
		return 0, ErrNegativeSqrt(guess)
	}

	for delta > Precision {
		old_guess := guess

		guess = Newton(guess, input)
		delta = math.Abs(guess - old_guess)
	}

	return guess, nil
}

func main() {
	fmt.Println(Sqrt(-2))
}
