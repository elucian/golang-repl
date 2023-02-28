package main

import (
	"fmt"
	"golang/algo"
)

func forth_menu() {
	fmt.Println("4:-------------------- ")
	fmt.Println("0 = clear ")
	fmt.Println("---------------------- ")
	fmt.Println("S = Sorting") 
	fmt.Println("---------------------- ")
	fmt.Println("ENTER => main menu")
	//main menu is in main.go
}

func forth() {
	var option = "_"
	clear()
	forth_menu()
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		switch option {
		case "0":
			clear()
    case "S":
      algo.TestSorting()
		default:
			clear(); print_menu()
			return
		} //switch
		option = "_"
	} // for
} // func
