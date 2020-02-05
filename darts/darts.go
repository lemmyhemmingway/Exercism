package darts

import "math"

func Score(x, y float64) int {
	r := math.Sqrt((x * x) + (y * y))
	if r <= 10 && r > 5 {
		return 1
	} else if r <= 5 && r > 1 {
		return 5
	} else if r <= 1 && r >= 0 {
		return 10
	} else {
		return 0
	}

}
