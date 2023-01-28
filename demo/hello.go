package demo

var message string = "Hello Wordl"

import (
	"fmt"
)

func Hello() {
	fmt.Println(message)
	fmt.Println('-' * 10)
	fmt.Println("Go \n" + 
	            "is \n" +
				"ossome \n")	
}
