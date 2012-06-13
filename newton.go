package main

import (
  "fmt"
  "math"
)

var tests = []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
const Threshold = 0.01

func Newton(z, x float64) float64 {
  return z - ((math.Pow(z, 2) - x) / (2 * z))
}

func Sqrt(input float64) float64 {
  for _, value := range tests {
    result := Newton(value, input)

    if delta := math.Abs(value - result); delta <= Threshold {
      return value
    } else {
      // fmt.Println(value, delta)
    }
  }

  return 0.0
}

func main() {
  fmt.Println(16, Sqrt(16))
  fmt.Println(25, Sqrt(25))
  fmt.Println(81, Sqrt(81))
}