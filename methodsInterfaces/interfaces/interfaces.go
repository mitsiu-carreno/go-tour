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