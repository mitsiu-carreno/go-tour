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