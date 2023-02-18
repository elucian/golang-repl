package funcs

import (
    "fmt"
    "math/rand"
)

func TestRandomize() {
    // seed a new source
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    for i:=0; i<10; i++ {    
      fmt.Print(r2.Intn(100), ",")
    }
    fmt.Println()
  
    // seed the same source
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    for i:=0; i<10; i++ {    
      fmt.Print(r3.Intn(100), ",")
    }
    fmt.Println()
}