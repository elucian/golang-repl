package main

/* ---------------------------------
   an application for learning Go.
   by Elucian Moise on Sage-Code
  --------------------------------- */

import (
	"fmt"
	demo "golang/demo"
	maps "golang/maps"
)

/* show on console the options to
   explain what the app does */
func first_menu() {
	fmt.Println("---------------------- ")
  fmt.Println("        FIRST          ")
	fmt.Println("---------------------- ")
	fmt.Println("M = Map demo")
	fmt.Println("I = Interpolation")
	fmt.Println("F = For while")
	fmt.Println("R = For Range")
	fmt.Println("S = Sqrt func")
  fmt.Println("E = Error handling")
	fmt.Println("--------------------- ")
	fmt.Println("ENTER => main menu")
  // next menu in in second.go
}


/* entry point, start the main loop */
func first() {
	option := "_"
	fmt.Println("Run demo: ")
	first_menu()
	// the actual loop (cycle)
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		// fmt.Print(option)
		if option == "Q" {
			break
		} else if option == "0" {
			clear()
		} else if option == "M" {
			maps.Init()
		} else if option == "I" {
			demo.ItpRun()
		} else if option == "F" {
			demo.ForWhile()
		} else if option == "E" {
			demo.TestErrors()
		} else if option == "S" {
			//sqrt_func.go file
			demo.Test_SQRT(9)
			demo.Test_SQRT(10)
			demo.Test_SQRT(11)
		} else if option == "R" {
			demo.Range()
		} else {
      clear(); print_menu()
      return;
		}
    option = "_"
	} //for
}
