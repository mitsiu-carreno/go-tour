package interfaces

import (
	"strings"
	"fmt"
	"math"
	"time"
	"io"
)

// MyFloat non-struct type
type MyFloat float64
// Vertex struct type
type Vertex struct{
	X, Y int
}
// Abser interface declaration with method Sign
type Abser interface{
	Sign() float64
}
// Sign method for Vertex pointer
func (v *Vertex)Sign()float64{
	if v.X < 0 {
		return -10
	}
	return 10
	
}
//Sign method but on MyFloat type
func (f MyFloat)Sign()float64{
	if f < 0 {
		return -5
	}
	return 5
	
}
// BasicInterface an interface type is defined as a set of method signature
func BasicInterface(){
	var a Abser
	var b Abser
	f := MyFloat(1)
	v := Vertex{2,3}

	a = f	// a MyFloat implements Abser
	b = &v 	// b Vertex implements Abser
	// a = v // as v is a Vertex (not *Vertex) a does not implement Abser (ERROR)

	fmt.Println(a.Sign())
	fmt.Println(b.Sign())
}

// T struct type
type T struct{
	S string
}
// I interface set M metod
type I interface{
	M()
}
// M This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so. 
func (t T)M(){
	fmt.Println(t.S)
}
// ImplementImplicit a type implements an interface by implementing its methods. Theres no explicit declaration (no "implement" key)
func ImplementImplicit(){
	// Implicit interfaces decouple the definition of an interface from its implementation,
	// Which could appear in any package withour prearrangement 
	var t I = T{"WORDS"}
	t.M()
}

// F type for example
type F float64
// M method with F signature
func (f F)M(){
	fmt.Println(f)
}
func describe(i I){
	fmt.Printf("(%v, %T)\n", i, i)
}
// InterfaceValues select a method from the interface depending on the type
func InterfaceValues(){
	var i I
	i = T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

// Interf custom interface
type Interf interface{
	handleNil()
}
// MyType custom struct
type MyType struct{
	S string
}
func (t *MyType)handleNil(){
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
// InterfaceNil if the concrete value inside an interface is nil, the method will be called with a nil receiver
func InterfaceNil(){
	// A nil interface value holds neither a value nor a concrete type 
	var i Interf
	// i.handleNil() // Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call
	var t *MyType 

	i = t
	fmt.Printf("(%v, %T)\n", i, i)
	// In some laguages this would trigger a null pointer exception, but in go it is common to write methods that gracefully handle being called with a nil receiver
	i.handleNil()

	i = &MyType{"Not nil anymore"}
	fmt.Printf("(%v, %T)\n", i, i)
	i.handleNil()
}

// EmptyInterface may hold values of any type (Every type implements at least zero methods)
func EmptyInterface(){
	// Empty interfaces are used by code that handles values of unknown type
	// For example "fmt.Print" takes any number of arguments of type "interface{}"
	var i interface{}
	
	fmt.Printf("(%v, %T)\n", i, i)
	
	i = 42
	fmt.Printf("(%v, %T)\n", i, i)

	i = "Hello"
	fmt.Printf("(%v, %T)\n", i, i)
}

// TypeAssertions provides access to an interface value's underlying concrete value "t := i.(T)"
func TypeAssertions(){
	// "t := i.(T)" asserts that the interface value "i" holds the concrete type "T" and assigns the underlying "T" value to the variable "t"
	
	var i interface{} = "Hello"

	// If "i" does not hold a "T", the statement will trigger a panic
	//f = i.(float64)

	// To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports wheter the assertion succeded (if not, prevents a panic) 
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

}

func getType(i interface{}){
	// The declaration in a type switch has the same syntax as a type assertion "i.(T)" but the specific 
	// type "T" is replaced with the keyword "type"
	switch v := i.(type){
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	// In the default case the variable v is of the same interface type and value as i
	default:
		fmt.Printf("%v for sure is neither an int nor a string\n", v)
	}
}
// TypeSwitch is a construct that permits several type assertions series
func TypeSwitch(){
	// A type switch is like a regular switch statement, but the cases in a type switch specify 
	// types (not values), and those values are compared against the type of the value held by the given 
	// interface value

	getType(20)
	getType("hello")
	getType(true)
}

// Person type for stringer example
type Person struct{
	Name string
	Age int
}
// String is the stringer an overrides the Println method below
func (p Person)String()string{
	return fmt.Sprintf("%v (%v years),", p.Name, p.Age)
}
// Stringers is one of the most ubiquitous interfaces is Stringer defined by the "fmt" package.
func Stringers(){
	a := Person{"Mitsiu", 24}
	b := Person{"Kary", 24}
	fmt.Println(a, b)
}

// MyError custom struct for errors
type MyError struct{
	When time.Time
	What string
}
// Error custom stringer
func (e *MyError)Error()string{
	return fmt.Sprintf("%s, at %v", e.What, e.When)
}
func run() error{
	return &MyError{
		time.Now(),
		"Triggered error",
	}
}
// Errors go programs express error state with "error value" 
func Errors(){
	// The error type is a built-in interface similar to fmt.Stringer
	if err := run(); err != nil{
		fmt.Println(err)
	}
}

// Readers the io.Reader package represents the read end of a stream of data
// the go standard library contains many implementations of these interfaces, including files, network connections, compressors, ciphers... (https://golang.org/search?q=Read#Global)
func Readers(){
	r := strings.NewReader("this is awesome")
	// Creates a strings.reader and consumes its output 8 bytes at the time
	b := make ([]byte, 8)

	for{
		// The io.Reader interface has a Read method " func(T) Read(b []bytes)(n int, err error)"
		// Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends
		n, err := r.Read(b)
		fmt.Printf("n = %v; err = %v; b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF{
			break
		}
	}
}