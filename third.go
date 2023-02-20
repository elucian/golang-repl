package main

import (
	"fmt"
	funcs "golang/funcs"
)

func third_menu() {
	fmt.Println("3:-------------------- ")
	fmt.Println("0 = clear ")
	fmt.Println("---------------------- ")
	fmt.Println("I = Interfsces")
	fmt.Println("---------------------- ")
	fmt.Println("ENTER => main menu")
	//main menu is in main.go
}

func third() {
	var option = "_"
	clear()
	third_menu()
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		switch option {
		case "0":
			clear()
    case "I":
      funcs.TestInterface()
		default:
			clear(); print_menu()
			return
		} //switch
		option = "_"
	} // for
} // func
