package demo

import (
	"fmt"
  "strings"
)

var message string = "Hello World!"

func Hello() {
	fmt.Println(message)
	fmt.Println( 
      strings.Repeat("-", len(message)))
	fmt.Println("Go \n" + 
	            "is \n" +
				"ossome \n")	
}
