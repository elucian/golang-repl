package main

import (
	"fmt"
	"golang/algo"
	"golang/str"
)

func forth_menu() {
	fmt.Println("---------------------- ")
	fmt.Println("         FORTH         ")
	fmt.Println("---------------------- ")
  fmt.Println("P = Parse String")
  fmt.Println("S = String Functions")
  fmt.Println("J = JSON Parsing")
  fmt.Println("F = Format string")
  fmt.Println("T = Text template")
  fmt.Println("R = Regular expressions") 
	fmt.Println("G = Sorting")   
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
    case "P":
      str.TestParser()
    case "J":
      str.TestJSON()
    case "S":
      str.TestFunctions()
    case "F":
      str.TestFormat()
    case "T":
      str.TestTemplate()
    case "R":
      str.TestRegex()
    case "G":
      algo.TestSorting()
		default:
			clear(); print_menu()
			return
		} //switch
		option = "_"
	} // for
} // func
