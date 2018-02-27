// Package moreTypes covers pointers, arrays, slices, maps
package moreTypes

import (
	"fmt"
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
	v := Vertex{1,4}
	v.X = 3
	fmt.Println(v)
	
}
// Vertex struct set a collection of fields or methods
type Vertex struct{
	X int 
	Y int
}

// LoopingArray when looping over an array, slice, string or map or reading from a channel use range
func LoopingArray(){
	//for 
}