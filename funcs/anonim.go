package funcs

import "fmt"

/* high order function that
   return anonymous function */
func intSeq() func() int {
    i := 0
    // create anonymous function
    return func() int {
        i++; return i
    }
}

func Anonymous() {
    // instantiate a closure
    first := intSeq()
    // instantiate second closure
    second  := intSeq()
    // execute closure 
    fmt.Println(first())
    fmt.Println(first())
    fmt.Println(second())
    fmt.Println(first()) 
    fmt.Println(second()) 
}