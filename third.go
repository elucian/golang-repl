package main

import (
	"fmt"
	// "golang/demo"
	"golang/funcs"
  "golang/chans"
	// structs "golang/structs"
)

func third_menu() {
	fmt.Println("3:-------------------- ")
	fmt.Println("0 = clear ")
	fmt.Println("---------------------- ")
	fmt.Println("I = Interfsces")
  fmt.Println("S = String functions")
  fmt.Println("C = Channels")  
  fmt.Println("B = Buffered channel")
  fmt.Println("W = Wait for goroutine")
  fmt.Println("G = Generic")  
  fmt.Println("E = Select channels")  
  fmt.Println("N = Closing channels") 
	fmt.Println("---------------------- ")
	fmt.Println("ENTER => main menu")
	//main menu is in main.go
}

func third() {
	var option = "_"
	clear()
	third_menu()
	for option != "Q" {
		fmt.Print(">>")
		fmt.Scanf("%s", &option)
		switch option {
		case "0":
			clear()
    case "I":
      funcs.TestInterface()
    case "S":
      funcs.TestString()
    case "C":
      chans.TestChannels()
    case "B":
      chans.TestBuffer()
    case "G":
      //structs.TestGeneric()
      fmt.Print("Not supported")
    case "E":
      chans.TestSelect()
    case "W":
      chans.ChanSync()
    case "N":
      chans.Closing()
		default:
			clear(); forth()
			return
		} //switch
		option = "_"
	} // for
} // func
