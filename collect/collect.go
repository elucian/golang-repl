package collect

import (
	"fmt"
)

type Person struct {  
    name string
    age  int
}

func Struct() {
    //creating struct data
    p1 := Person{
        name: "Liviu",
        age: 55,
    }
    fmt.Println("Person 1", p1)
    //creating data 
    p2 := Person{"Elucian", 56}
    fmt.Println("Person 2", p2)
}
