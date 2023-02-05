package demo

import (
	"fmt"
	"math"
)

var lim float64 = 8

// limited power
func pow_lim(x, n float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Local_scope() {
	fmt.Println(pow_lim(3, 3))
  lim = 20
  fmt.Println(pow_lim(3, 3))
}