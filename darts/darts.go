package darts

import "math"

func Score(x, y float64) int {
	radius := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	if radius > 10.0 {
		return 0
	} else if radius <= 10 && radius > 5 {
		return 1
	} else if radius <= 5 && radius > 1.0 {
		return 5
	} else {
		return 10
	}
}
