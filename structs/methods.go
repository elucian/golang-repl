package structs
/* Go supports _methods_ defined on struct types.  Methods have a special parameter, that is called a receiver. The receiver may be a pointer or a value based.
*/
import "fmt"

type rect struct {
	width, height int
}

// This method has a "receiver type" pointer "r *rect"
func (r *rect) area() int {
	return r.width * r.height
}

// Here's an example of a "value receiver": "r rect"
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func TestMethods() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

  // Here we test call using a pointer qualifier "rp"
	rp := &r // prepare a pointer qualifier for r instance
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
/* NOTES:
Go automatically handles conversion between values and pointers for a method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct. 
*/