package geometry

import (
	"math"
)

func Area(length, width float64) float64 {
	return length * width
}

// Function that can be exported
func Diagonal(length, width float64) float64 {
	return math.Sqrt((length * length) + (width * width))
}
