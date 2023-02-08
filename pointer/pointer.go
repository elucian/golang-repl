package pointer

import "fmt"

//no pointer (input only)
func zeroval(ival int) {
	ival = 0 // try to modify
}

//pointer (input & output)
func zeroptr(iptr *int) {
	*iptr = 0 // modify value
}

func TestPointer() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i) //1
	zeroptr(&i)
	fmt.Println("zeroptr:", i) //0
	fmt.Println("pointer:", &i)
}
