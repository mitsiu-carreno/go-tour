package interfaces

import (
	"fmt"
	"math"
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