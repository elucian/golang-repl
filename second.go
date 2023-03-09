package main

import (
	"fmt"
	collect "golang/collect"
	funcs   "golang/funcs"
	pointer "golang/pointer"
	structs "golang/structs"
)

func second_menu() {
	fmt.Println("---------------------- ")
	fmt.Println("        SECOND         ")
	fmt.Println("---------------------- ")
	fmt.Println("C = Closure")
	fmt.Println("P = Pointer")
	fmt.Println("A = Anonymous function")
	fmt.Println("H = Methods & Interface ")
	fmt.Println("L = Linked list")  
	fmt.Println("M = Map literals")
	fmt.Println("S = Struct test")
 	fmt.Println("E = Embedded struct") 
	fmt.Println("F = Factorial")
	fmt.Println("W = Swap values")
	fmt.Println("V = Variadic functions")
	fmt.Println("R = Recover from panic")
	fmt.Println("T = Timer")
	fmt.Println("G = Goroutines")
  fmt.Println("Z = Randomize")
	fmt.Println("---------------------- ")
	fmt.Println("ENTER => main menu")
	//main menu is in main.go
}

// start second selection
func second() {
	var option = "_"
	clear()
	second_menu()
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		switch option {
		case "0":
			clear()
		case "H":
			//methods.go
			structs.TestMethods()
		case "A":
			funcs.Anonymous()
		case "S":
			structs.TestStruct()
    case "E":
      structs.TestEmbedded()
		case "W":
			funcs.TestSwap()
		case "F":
			funcs.TestFactorial(5)
			funcs.TestFactorial(6)
			funcs.TestFactorial(7)
		case "C":
			funcs.TestClosure(2)
		case "P":
			pointer.TestPointer()
		case "V":
			funcs.TestVariadic()
		case "L":
			structs.LinkedList()
		case "M":
			collect.MapLiterals()
		case "R":
			funcs.TestRecover()
		case "T":
			funcs.TestTimer()
    case "G": 
      funcs.TestGoroutine()
    case "Z":
      funcs.TestRandomize()
		default:
			clear(); print_menu()
			return
		} //switch
		option = "_"
	} // for
} // func
