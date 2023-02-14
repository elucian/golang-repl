package main

/* ---------------------------------
   golang-repl
   an application for learning Go.
   by Elucian Moise on Sage-Code
  --------------------------------- */

import (
	"fmt"
	array "golang/array"
	demo "golang/demo"
	maps "golang/maps"
)

/* show on console the options to 
   explain what the app does */
func print_menu() {
	fmt.Println("---------------------- ")
  fmt.Println("0 = clear ")
	fmt.Println("Q = quit ")
	fmt.Println("1 = Hello World")
	fmt.Println("2 = For loop")
	fmt.Println("3 = Ladder")
	fmt.Println("4 = Infinite loop")
	fmt.Println("5 = Value switch")
	fmt.Println("6 = Cond switch")
	fmt.Println("7 = Local scope")
	fmt.Println("8 = Array Init")
	fmt.Println("9 = Array Slice")
	fmt.Println("D = Map demo")
	fmt.Println("B = Interpolation")
	fmt.Println("W = For while")
	fmt.Println("R = For Range")
	fmt.Println("T = Sqrt func")  
	fmt.Println("--------------------- ")
  fmt.Println("X = neXt screen")
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
    } else if option == "1" {
			demo.Hello()
		} else if option == "2" {
			demo.For_loop()
		} else if option == "3" {
			demo.Ladder()
		} else if option == "4" {
			demo.Infinite_loop()
		} else if option == "5" {
			demo.Value_switch()
		} else if option == "6" {
			demo.Cond_switch()
		} else if option == "7" {
			demo.Local_scope()
		} else if option == "8" {
			array.Init()
		} else if option == "9" {
			array.Slice()  
		} else if option == "D" {
			maps.Init()
		} else if option == "B" {
			demo.ItpRun()
		} else if option == "W" {
			demo.ForWhile()
		} else if option == "T" {
      //sqrt_func.go file
			demo.Test_SQRT(9)  
			demo.Test_SQRT(10)
			demo.Test_SQRT(11)
		} else if option == "R" {
			demo.Range()
    } else if option == "X" {
      select_secod()
    } else {
			fmt.Println("invalid option")
			print_menu()
		}
	}
}
