package demo

import (
    "fmt"
    "os"
)

func DemoExit() {
    // expect ths to print
    fmt.Println("exiting ...")

    // is this going to be print?
    defer fmt.Println("done !")

    os.Exit(3) // error
}