package main

import (
	"fmt"
	"math"
	"os"
	"time"
	"net/http"
)

type Vertex struct {
	X, Y float64
}

/*
   The method receiver appears in its own argument list between the func keyword and the method name.
*/
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

/*
   You can define a method on any type you define in your package, not just structs.
   You cannot define a method on a type from another package, or on a basic type.
*/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func interfaces() {
	type Abser interface {
		Abs() float64
	}

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// This causes the  code not to compile.
	//a = v

	fmt.Println(a.Abs())

	/*
		A type implements an interface by implementing the methods.
		There is no explicit declaration of intent.
	*/

	type Reader interface {
		Read(b []byte) (n int, err error)
	}

	type Writer interface {
		Write(b []byte) (n int, err error)
	}

	type ReadWriter interface {
		Reader
		Writer
	}

	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}

func methods() {
	/*
		There are two reasons to use a pointer receiver.
		First, to avoid copying the value on each method call (more efficient if the value type is a large struct).
		Second, so that the method can modify the value that its receiver points to.

	*/
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(MyFloat(-math.Sqrt2).Abs())
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

/*
	An error is anything that can describe itself as an error string.
	The idea is captured by the predefined, built-in interface type, error, with its single method, Error,
	returning a string.
*/
func errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.
}

func main() {
	methods()
	interfaces()
	errors()
}
