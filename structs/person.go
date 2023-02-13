package structs

import "fmt"

// define a stucture with two fields
// fields are not comma separated
type Person struct {
	name string
	age  int
}

//constructor
func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func TestStruct() {
	// ceate a person and ptint
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))
	// create n object "s"
	s := Person{name: "Sean", age: 50}
	fmt.Println("name=", s.name)
	// make a pointer to object
	sp := &s
	fmt.Println("age=", sp.age)
	//modify one filed
	sp.age = 51
	fmt.Println("age=", sp.age)
}
