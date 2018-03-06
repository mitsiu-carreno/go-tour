package interfaces

import (
	"fmt"
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