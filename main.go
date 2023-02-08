//golang-repl
package main

import (
	"fmt"
	array "golang/array"
	collect "golang/collect"
	demo "golang/demo"
	funcs "golang/funcs"
	maps "golang/maps"
	pointer "golang/pointer"
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
  fmt.Println("M = Map literals")
	fmt.Println("B = Interpolation")
  fmt.Println("C = Closure")
  fmt.Println("E = Factorial")
	fmt.Println("F = For while")
  fmt.Println("P = Pointer")
	fmt.Println("R = For Range")
	fmt.Println("S = Struct test")
	fmt.Println("T = Sqrt func")
  fmt.Println("W = Swap values")
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
		} else if option == "F" {
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
    } else if option == "W" {
      funcs.TestSwap()
    } else if option == "E" {
      funcs.TestFactorial(5)
      funcs.TestFactorial(6)
      funcs.TestFactorial(7)
    } else if option == "C" {
      funcs.TestClosure(2)
    } else if option == "P" {
      pointer.TestPointer()
    } else {
			fmt.Println("invalid option")
			print_menu()
		}
	}
}
