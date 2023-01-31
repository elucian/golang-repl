package array

import "fmt"

func Slice() {
  //declare an array
	primes := [6]int{2, 3, 5, 7, 11, 13}
  //making a slice
	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]
  fmt.Println(s[0]) // [3]
  fmt.Println(s[2]) // [7]
}