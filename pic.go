package main

import "code.google.com/p/go-tour/pic"
import "math/rand"

func Pic(x, y int) [][]uint8 {
	picture := make([][]uint8, y)

	for row := range picture {
		picture[row] = make([]uint8, x)

		for col := range picture[row] {
			// picture[row][col] = uint8(row^col + (row + col) / 2)
			// picture[row][col] = uint8(row ^ col + row * math.Pi)
			// picture[row][col] = uint8(row * col)
			picture[row][col] = uint8(rand.Intn(255))
			rand.Int()
			// picture[row][col] = uint8((row + col) / 2)
			// picture[row][col] = uint8(row^2 + col)
		}
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
