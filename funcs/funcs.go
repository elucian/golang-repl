package funcs

import "fmt"

func swap(a, b int) (x, y int) {
	x = b
	y = a
	return
}

func TestSwap() {
	a, b := swap(10, 11)
  fmt.Println("funcs.TestSwap()")
  fmt.Println("a = ", a)
  fmt.Println("b = ", b)
}
