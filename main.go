package main

/* ---------------------------------
   an application for learning Go.
   by Elucian Moise on Sage-Code
  --------------------------------- */

import (
	"fmt"
	array "golang/array"
	demo "golang/demo"
)

/* show on console the options to
   explain what the app does */
func print_menu() {
	fmt.Println("---------------------- ")
	fmt.Println("         MAIN          ")
	fmt.Println("---------------------- ")
	fmt.Println("Q = quit")
	fmt.Println("H = Hello World")
	fmt.Println("F = For loop")
	fmt.Println("L = Ladder")
	fmt.Println("E = Infinite loop")
	fmt.Println("V = Value switch")
	fmt.Println("G = Cond switch")
	fmt.Println("C = Local scope")
	fmt.Println("I = Array Init")
	fmt.Println("S = Array Slice")
	fmt.Println("I = Interpolation")
	fmt.Println("F = For while")
  fmt.Println("X = Exit")
	fmt.Println("--------------------- ")
	fmt.Println("SUBMENU: 0,1,2,3,4")
	// next menu in in second.go
}

/* clear the screen */
func clear() {
	fmt.Print("\033[H\033[2J")
}

/* entry point, start the main loop */
func main() {
	option := "_"
	fmt.Println("Run demo: ")
	print_menu()
	// the actual loop (cycle)
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		// fmt.Print(option)
		if option == "Q" {
			break
		} else if option == "0" {
			clear()
		} else if option == "H" {
			demo.Hello()
		} else if option == "L" {
			demo.Ladder()
		} else if option == "E" {
			demo.Infinite_loop()
		} else if option == "V" {
			demo.Value_switch()
		} else if option == "G" {
			demo.Cond_switch()
		} else if option == "C" {
			demo.Local_scope()
  	} else if option == "S" {
			array.Slice()
		} else if option == "I" {
			array.Init()
		} else if option == "F" {
			demo.For_loop()
    } else if option == "X" {
      demo.DemoExit();
		} else if option == "1" {
			clear()
			first()
		} else if option == "2" {
			clear()
			second()
		} else if option == "3" {
			clear()
			third()
		} else if option == "4" {
			clear()
			forth()
		} else {
			clear()
			print_menu()
		}
		option = "_"
	} //for
}
