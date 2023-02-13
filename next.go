package main

import (
	"fmt"
	structs "golang/structs"
)

func more() {
	fmt.Println("---------------------- ")
	fmt.Println("B = back ")
	fmt.Println("0 = clear ")
	fmt.Println("1 = methods ")
	fmt.Println("---------------------- ")
}

func select_secod() {
	var option = "_"
	clear()
	more()
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		switch option {
		case "0":
			clear()
		case "B":
			return
		case "1":
      //methods.go
			structs.TestMethods()  
		default:
			clear()
			fmt.Printf("Invalid %s.", option)
			more()
		} //switch
	} // for
} // func
