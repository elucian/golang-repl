package main

import (
	"fmt"
	"golang/algo"
	"golang/str"
)

func forth_menu() {
	fmt.Println("4:-------------------- ")
	fmt.Println("0 = clear ")
	fmt.Println("---------------------- ")
	fmt.Println("S = Sorting") 
  fmt.Println("P = Parse String")
  fmt.Println("F = String Functions")
  fmt.Println("J = JSON Parsing")
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
    case "P":
      str.TestParser()
    case "J":
      str.TestJSON()
    case "F":
      str.TestFunctions()
		default:
			clear(); print_menu()
			return
		} //switch
		option = "_"
	} // for
} // func
