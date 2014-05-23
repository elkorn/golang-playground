package main

/*
   Go's basic types are:
   bool
   string
   int  int8  int16  int32  int64
   uint uint8 uint16 uint32 uint64 uintptr
   byte // alias for uint8
   rune // alias for int32
        // represents a Unicode code point
   float32 float6k4
   complex64 complex128
*/

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"time"
)

var i, j int = 1, 2
var c, python, java = true, false, "no!"

const f = "%T(%v)\n" // %Type and %value

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/*
   Numeric constants are high-precision values.
   An untyped constant takes the type needed by its context.
*/
const (
	Big   = 1 << 100
	Small = Big >> 99
)

// A struct is a collection of fields.
// (And a type declaration does what you'd expect.)
type Vertex struct {
	X int
	Y int
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// Due to the else, v can't be used here.
	return lim
}

func basicTypes() {
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

func typeConversions() {
	var x, y int = 3, 4
	ff := math.Sqrt(float64(3*3 + 4*4))
	// The same as:
	// var f float64 = math.Sqrt(3*3 + 4*4)
	var zz int = int(ff)
	fmt.Println(x, y, zz)
}

func constants() {
	xx := needInt(Small)
	fmt.Printf(f, xx, xx)
	xxx := needFloat(Small)
	fmt.Printf(f, xxx, xxx)
	xxx = needFloat(Big)
	fmt.Printf(f, xxx, xxx)
	// fmt.Printf(f, needInt(Big)) // Causes an overflow.
}

func loops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func ifStatements() {
	fmt.Println(
		sqrt(2),
		sqrt(-4),
	)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func structs() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
}

func pointers() {
	/*
		Go has pointers, but no pointer arithmetic.
		Struct fields can be accessed through a struct pointer.
		The indirection through the pointer is transparent.
	*/
	p := Vertex{6, 7}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}

func structLiterals() {
	/*
		A struct literal denotes a newly allocated struct value by listing the values of its fields.
		You can list just a subset of fields by using the Name: syntax.
		(And the order of named fields is irrelevant.)
		The special prefix & constructs a pointer to a newly allocated struct.
	*/
	var (
		p = Vertex{1, 2}  // has type Vertex
		q = &Vertex{1, 2} // has type *Vertex
		r = Vertex{X: 1}  // Y:0 is implicit
		s = Vertex{}      // X:0 and Y:0
	)

	fmt.Println(p, q, r, s)
}

func theNewFunction() {
	v := new(Vertex) // Returns a ptr to a newly allocated, zeroed value.
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
	// Equivalent to:
	var t *Vertex = new(Vertex)
	fmt.Println(t)
	t.X, t.Y = 11, 9
	fmt.Println(t)
}

func arrays() {
	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
}

func slices() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

func slicingSlices() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)
	fmt.Println("p[1:4] ==", p[1:4])
	fmt.Println("p[:3] ==", p[:3])
	fmt.Println("p[4:] ==", p[4:])
}

func makingSlices() {
	a := make([]int, 5) // 5-element slice of an underlying zeroed anonymous array.
	printSlice("a", a)
	b := make([]int, 0, 5) // An empty slice with cap(5).
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5] // len(d)=3, cap(d)=3
	printSlice("d", d)
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("z is nil!")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func ranges() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d == %d\n", i, v)
	}

	for i := range pow {
		// skipping the second value
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	for _, v := range pow {
		// skipping the first value
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func maps() {
	type Vertex struct {
		Lat, Long float64
	}

	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)

	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

func mutatingMaps() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 52
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot)
	fmt.Println(hypot(3, 4))
}

func functionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func switches() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func switchEvalOrder() {
	// Switches are evaluated from top to bottom until a successful case is hit.
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func unconditionalSwitch() {
	t := time.Now()
	// This is the same as `switch true`.
	// It can be useful for substituting long if-else chains.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	fmt.Print(i, j, c, python, java)

	basicTypes()
	typeConversions()
	constants()
	loops()
	ifStatements()
	structs()
	pointers()
	structLiterals()
	theNewFunction()
	arrays()
	slices()
	slicingSlices()
	makingSlices()
	ranges()
	maps()
	mapLiterals()
	mutatingMaps()
	functionValues()
	functionClosures()
	switches()
	switchEvalOrder()
	unconditionalSwitch()
}
