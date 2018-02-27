// Package moreTypes covers pointers, arrays, slices, maps
package moreTypes

import (
	"fmt"
)

// Vertex struct set a collection of fields or methods
type Vertex struct{
	X int 
	Y int
}

// Struct literal denotes a newly allocated struct value by listing the values of its fields
var (
	v1 = Vertex{}		// X:0 & Y:0 	
	v2 = Vertex{Y:1}	// X:0 implicitly
	v3 = Vertex{2,3}	// Type Vertex
	v4 = &Vertex{4,5}	// Type *Vertex

)

// Pointers hold the memory address of a value 
func Pointers (){
	i, j := 42, 1052

	// The "&" operator generates a pointer to its operand
	p := &i	//Point to i	
	fmt.Printf("i = %v, p = %v\n", i, p)
	// The type "*T" is a pointer to a T value. Its zero value is nil 
	*p = 21	// Set i through the pointer
	fmt.Printf("i = %v, p = %v, *p = %v\n", i, p, *p)

	p = &j	//Point to j
	fmt.Printf("i = %v, p = %v, *p = %v\n", i, p, *p)
	*p = *p / 2	// Divide j through the pointer
	fmt.Printf("i = %v, p = %v, *p = %v, j = %v\n", i, p, *p, j)

}

// Structs are collections of fields
func Structs(){
	fmt.Println(Vertex{1,2})

	// Struct fields are accessed using a dot.
	v := Vertex{X:1,Y:4}
	v.X = 3
	fmt.Println(v)
	
	//Pointers to struct; Struct fields can be accessed through a struct pointer.
	p := &v
	p.X = 1e9	// or (*p).X is also possible
	fmt.Println(v)

	//Struct literals
	fmt.Println(v1, v2, v3, v4)
}

// ArrayStruct shows you how to create a basic array
func ArrayStruct(){
	// The type "[n]T" is an array of n values of type T 
	// An arrays length its part of its type, so ARRAYS CANNOT BE RESIZED
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0])
	fmt.Println(a)
}

// SliceStruct
func SliceStruct(){
	// While arrays have fixed sizes. An
}
// LoopingArray when looping over an array, slice, string or map or reading from a channel use range
func LoopingArray(){
	//for 
}