package main

import (
	"math"
)

func precisionRound(x float64, decimalPlaces int) float64 {
	n := math.Pow10(decimalPlaces)
	return math.Round(x*n) / float64(n)
}

func sign(x float64) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}
