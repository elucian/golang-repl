package maps

import "fmt"

var m map[string] int

func Init() {
	m = make(map[string] int)
	m["test"] = 1
  m["demo"] = 2
  m["collect"] = 3
  fmt.Println(m)
}