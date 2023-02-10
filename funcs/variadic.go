package funcs

import "fmt"

/* ... is a prefix. part of syntax
   indicates: variable arguments 
   it can be an array or a slice */
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    // read values, ignore the index
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func TestVariadic() {
    // various calls with 2,3,4 arguments
    sum(1, 2)
    sum(1, 2, 3)
    // expand the the array and send arguments
    // using spreading operator ... creazy!
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}