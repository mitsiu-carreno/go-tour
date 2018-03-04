package methods

import(
	"fmt"
)

// Vertex struct is used to bind the method Sum
type Vertex struct{
	X, Y float64
}
// Sum have a special reciever argument specifing the type "Vertex" named "v"
func (v Vertex) Sum ()float64{
	return v.X + v.Y
}
// BasicMethod shows the basic method calling 
func BasicMethod(){
	// Go doesn't have classes. However you can define methods on types 
	v := Vertex{3, 4}
	fmt.Println(v.Sum())
}

// MyString is a simple non struct type
type MyString string
// Concat simple method to add " for real" to strings
func (s MyString)Concat()MyString{
	return s + " for real"
}

// NonStruct is an example of a method of non struct type
func NonStruct(){
	s := MyString("I love you")
	fmt.Println(s.Concat())
}

// PointerScale and ValueScale perform the exact same operation but one receive a pointer other the value
func (v *Vertex) PointerScale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}
// ValueScale and PointerScale perform the exact same operation but one receive a pointer other the value
func (v Vertex) ValueScale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}
// Scale performs exactly as PointerScale but isn´t a method
func Scale(v *Vertex, f float64){
	v.X = v.X * f
	v.Y = v.Y *f
}
// PointerReceiver means that the receiver type has the literal sintax *T for some type T (T cannot itself be a pointer)
func PointerReceiver(){
	pointerReceiver := Vertex{3,4}
	valReceiver := Vertex{3,4}
	normalFunction := Vertex{3,4}

	// Since methods often need to modify their receiver, pointer receivers are more common than value receivers
	pointerReceiver.PointerScale(10)
	// methods with pointer receiver takes either a value or a pointer as the receiver when they are called ()
	p := &pointerReceiver
	p.PointerScale(2)
	fmt.Println(p)

	// Value receiver operates on a copy of the original Vertex value 
	valReceiver.ValueScale(10)
	// The Scale method must have a pointer receiver to change the vertex value declared in this function

	Scale(&normalFunction, 10)

	fmt.Println("pointer receiver, modified", pointerReceiver)
	fmt.Println("value receiver, not modified", valReceiver)
	fmt.Println("not a method but pointer receiver, modified", normalFunction)

	// In general all methods on a given type should have either value or ponter receivers, but not a mixture of both 
}

