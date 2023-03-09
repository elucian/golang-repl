package funcs

import (
	"fmt"
)

// fast integer power a^n
func pwr(a, n int) int {
    ans := 1; // final answer
    last_bit := 0;
    for (n > 0) {
        last_bit = n & 1;
        // Check if last bit is set
        if (last_bit == 1) {
            ans = ans * a;
        }
        a = a * a;
        n = n >> 1;
    }
    return ans;
}

// absoulte
func abs(x float64) float64 {
  if x < 0 { 
     return -x
  }
  return x
}

func Sqrt(x float64, p int) float64 {
   z := 1.0            // result  
   n := pwr(10, p)     // 10^p
   m := 1.0/float64(n) // precision
   for abs(z*z - x) > m  {
      z -= (z*z - x) / (2*z) 
   }
   return z
}

func Test_SQRT(x float64) {
  fmt.Printf("sqrt(%v,2)  = %v \n",x,Sqrt(x, 2));
	fmt.Printf("sqrt(%v,5)  = %v \n",x,Sqrt(x, 5));
  fmt.Printf("sqrt(%v,10) = %v \n",x,Sqrt(x,10));
}