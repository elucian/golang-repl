package funcs

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestClosure(p int) {
	var f = fib()
	for x := f(); x < 20; x = f() {
		if x < p {
			continue
		}
		fmt.Println(x)
	}
}