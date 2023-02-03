//golang-repl
package main

import (
	"fmt"
	array "golang/array"
	demo "golang/demo"
	maps "golang/maps"
  collect "golang/collect"
  funcs "golang/funcs"
)

func print_menu() {
	fmt.Println("----------------- ")
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
	fmt.Println("A = Map demo")
	fmt.Println("B = Interpolation")
	fmt.Println("C = For while")
	fmt.Println("T = Sqrt func")
	fmt.Println("R = For Range")
	fmt.Println("S = Struct test")
  fmt.Println("M = Map literals")
  fmt.Println("F = Muliple results")
	fmt.Println("----------------- ")
}

func main() {
	option := "_"
	fmt.Println("Run demo: ")
	print_menu()
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		// fmt.Print(option)
		if option == "Q" {
			break
		} else if option == "0" {
      fmt.Print("\033[H\033[2J")
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
		} else if option == "A" {
			maps.Init()
		} else if option == "B" {
			demo.ItpRun()
		} else if option == "C" {
			demo.ForWhile()
		} else if option == "T" {
			demo.Test_SQRT(9)
			demo.Test_SQRT(10)
			demo.Test_SQRT(11)
		} else if option == "R" {
			demo.Range()
    } else if option == "S" {
      collect.Struct()
    } else if option == "M" {
      collect.MapLiterals()
    } else if option == "F" {
      funcs.TestSwap()
		} else {
			fmt.Println("invalid option")
			print_menu()
		}
	}
}
