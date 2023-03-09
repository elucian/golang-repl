package demo

import (
  "os"
  "fmt"
)

func TestPanic() {
   // handle error using panic
   _, err := os.Create("tmp/file")
    // conditional panic
   if err != nil {
       panic(err)
   }
   fmt.Println("done")
}