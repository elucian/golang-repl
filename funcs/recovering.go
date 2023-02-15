/* Go makes it possible to _recover_ from a panic, by using the `recover` built-in function. A `recover` can stop a `panic` from aborting the program and let it continue with execution instead.*/

/* An example of where this can be useful: a server wouldn't want to crash if one of the client connections exhibits a critical error. Instead, the server would want to close that connection and continue serving other clients. In fact, this is what Go's `net/http` does by default for HTTP servers.*/

package funcs

import "fmt"

/* recover() must be called within a deferred function. When the enclosing function panics, the defer will activate and a recover() call within it will catch the panic error message and report it.*/

func TestRecover() {
	fmt.Println("protect this function")
	defer func() {
		if e := recover(); e != nil {
			// The return value of recover() is the error
			fmt.Println("Recovered from: ", e)
		}
	}()
	fmt.Println("get ready for panic")
	// create an error on purpose
	fmt.Println("paniking ...")
	panic("My error")

	/*Next print will not run, because "mayPanic" panics. The execution of "main" stops at the point of the  panic and resumes in the deferred closure.*/

	fmt.Println("This is not executed.")
}
