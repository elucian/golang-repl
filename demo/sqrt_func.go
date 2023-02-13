package demo

import (
	"fmt"
)

func Sqrt(x float64, p int) float64 {
   z := 1.0;
   for i :=1; i < p; i++ {
      z -= (z*z - x) / (2*z) 
   }
   return z
}

func Test_SQRT(x float64) {
	fmt.Printf("sqrt(%v) = %v \n",x,Sqrt(x, 10));
}