package demo

import "fmt"

//declare a slice of 7 elements
var pow = []int{1, 2, 4, 8, 16, 32, 64}

func Range() {
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
}
