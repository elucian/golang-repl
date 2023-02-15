package main

import (
	"fmt"
	collect "golang/collect"
	funcs "golang/funcs"
	pointer "golang/pointer"
	structs "golang/structs"
)

func more() {
	fmt.Println("---------------------- ")
	fmt.Println("0 = clear ")
	fmt.Println("---------------------- ")
	fmt.Println("C = Closure")
	fmt.Println("P = Pointer")
	fmt.Println("A = Anonymous function")
	fmt.Println("M = Methods & Interface ")
	fmt.Println("L = Map literals")
	fmt.Println("S = Struct test")
	fmt.Println("F = Factorial")
	fmt.Println("W = Swap values")
	fmt.Println("V = Variadic functions")
	fmt.Println("R = Recover from panic")
	fmt.Println("T = Timer")
	fmt.Println("G = Goroutines")
	fmt.Println("---------------------- ")
	fmt.Println("ENTER = main menu")
	//main menu is in main.go
}

func secod() {
	var option = "_"
	clear()
	more()
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		switch option {
		case "0":
			clear()
		case "M":
			//methods.go
			structs.TestMethods()
		case "A":
			funcs.Anonymous()
		case "S":
			structs.TestStruct()
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
			collect.MapLiterals()
		case "R":
			funcs.TestRecover()
		case "T":
			funcs.TestTimer()
    case "G": 
      funcs.TestGoroutine()
		default:
			clear()
			print_menu()
			return
		} //switch
		option = "_"
	} // for
} // func
