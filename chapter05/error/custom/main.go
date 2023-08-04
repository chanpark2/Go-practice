package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. 입력값: %f", n)
	}

	return math.Sqrt(n), nil
}
