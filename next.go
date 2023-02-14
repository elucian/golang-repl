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
	fmt.Println("B = back ")
	fmt.Println("C = Closure")
	fmt.Println("P = Pointer")
	fmt.Println("F = Anonymous Function")
	fmt.Println("H = methods ")
	fmt.Println("M = Map literals")
	fmt.Println("S = Struct test")
	fmt.Println("E = Factorial")
	fmt.Println("W = Swap values")
	fmt.Println("V = Variadic functions")
	fmt.Println("---------------------- ")
}

func select_secod() {
	var option = "_"
	clear()
	more()
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		switch option {
		case "0":
			clear()
		case "B":
			clear()
			print_menu()
			return
		case "H":
			//methods.go
			structs.TestMethods()
		case "F":
			funcs.Anonymous()
		case "S":
			structs.TestStruct()
		case "W":
			funcs.TestSwap()
		case "E":
			funcs.TestFactorial(5)
			funcs.TestFactorial(6)
			funcs.TestFactorial(7)
		case "C":
			funcs.TestClosure(2)
		case "P":
			pointer.TestPointer()
		case "V":
			funcs.TestVariadic()
		case "M":
			collect.MapLiterals()
		default:
			clear()
			fmt.Printf("Invalid option \"%s\". \n", option)
			more()
		} //switch
	} // for
} // func
