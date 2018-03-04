// Package moreTypes covers pointers, arrays, slices, maps
package moreTypes

import (
	"strings"
	"fmt"
	"strconv"
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

	// This is an array literal
	var b = [3]bool{true, true, false}
	fmt.Println("Array literal (var b = [3]bool{true, true, false})")
	fmt.Println(b)
}

// SliceStruct While arrays have fixed sizes. A slice is dinamically-sized  SLICES ARE LIEK REFERENCES TO ARRAYS
func SliceStruct(){
	// The type []T is a slice with elements of type T
	// A slice is formed by specifing two indices, a low and a high bound, separated by a colon 
	// a[low: high] (Default = [0:length])
	// This selects a half-open range which includes the first element but excludes the second one
	
	primes := [6]int{2,3,5,7,11,13}	// This is an array
	var s []int = primes[1:4] 
	fmt.Printf("Array = %v, slice[1:4] = %v \n", primes, s)
	
	//-------------------------------------------IMPORTANT------------------------------------------//
	// A SLICE DOES NOT STORE ANY DATA it just describes a section of an underlying array			//
	// Changing elements of a slice modifies the corresponding elements of its underlying array		//
	// Other slices that share the same underlying array will see those changes 					//
	//-------------------------------------------IMPORTANT------------------------------------------//

	numbers := []int{0,1,2,3,4,5}
	a := numbers[1:4]
	b := numbers[:2]
	c := numbers[1:]
	fmt.Println("Original array: ", numbers)
	fmt.Println("slice[1:4]", a)
	fmt.Println("slice[:2]", b)
	fmt.Println("slice[1:]", c)


	// Slice literals are like an array literal without the length
	q := []int{2,5,3,8}
	r := []bool{true, false, true, false}
	t := []struct{
		int
		bool
	}{
		{5, true},
		{3, false},
		{10, false},
	}
	fmt.Println("slice literals (q := []int{2,5,3,8}), example:")
	fmt.Println(q, " // ",  r, " // ", t)

	// Slice has both length and capacity
	// The length "len(s)" of a slice is the number of elements it contains
	// The capacity "cap(s)" of a slice is the number of elements in the underlying array, counting from the first element in the slice
	y := []int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("len=%d, cap = %d, %v\n", len(y), cap(y), y)
	y = y[:0]
	fmt.Printf("len=%d, cap = %d, %v\n", len(y), cap(y), y)
	// You can extend a slice's length by re-slicing it
	y = y[:4]
	fmt.Printf("len=%d, cap = %d, %v\n", len(y), cap(y), y)
	y = y[2:]
	fmt.Printf("len=%d, cap = %d, %v\n", len(y), cap(y), y)
	y = y[1:]
	fmt.Printf("len=%d, cap = %d, %v\n", len(y), cap(y), y)

	// Nil slices
	// A zero value of slice is nil
	// A nil slice has zero length and capacity and no underlying array
	var z []int
	fmt.Printf("len=%d, cap = %d, %v\n", len(z), cap(z), z)

	// Create slice with make
	// The make function allocates a zeroed value and returns a slice that refers to that array
	// a := make([]int, 5) // len(a) = 5
	// To specify a capacity pass a third argument
	rt := make([]int, 5, 10) 
	fmt.Printf("make : len=%d, cap = %d, %v\n", len(rt), cap(rt), rt)

	// Slices of slices
	board := [][]string{
		[]string{"-","-","-"},
		[]string{"-","-","-"},
		[]string{"-","-","-"},
	}
	// board[xAxis][yAxis]
	board[0][0] = "X"
	board[0][2] = "X"
	board[1][2] = "O"

	for i := 0; i < len(board); i++{
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to slide
	k := []int{1,2,3,4,5}

	k = append(k, 6)
	fmt.Println("Append 6 to slice", k)


}
// LoopingArray when looping over an array, slice, string or map or reading from a channel use range
func LoopingArray(){
	pow := []int{1,2,4,8,16,32,64,128,256,512,1024}

	for i, v := range pow{
		fmt.Printf("2**%v = %v\n", i,v)
	}

	// Or drop the key
	for _, v := range pow{
		fmt.Println("Drop key '_'", v)
	}
}


// MapsStruct is a struct which maps keys to values (like JSON's)
func MapsStruct(){
	// The zero value of a map is "nil". A "nil" map has no keys, nor keys can be added 
	var m map[string]Vertex
	// The make function returns a map of a given type, initialized and ready to use
	m = make(map[string]Vertex)
	m["Key1"] = Vertex{
		1,2,
	}
	fmt.Println(m)
	fmt.Println(m["Key1"])

	// If the top-level type is just a type name, you can omit it from the elements of the literal 
	var t = map[string]Vertex{
		"test" : {10,12},
		"test2" : {12,13},
	}
	fmt.Println(t)

	// Insert or update an element in map m[key] = elem
	t["test"] = Vertex{10, 11}
	t["test3"] = Vertex{14, 15}
	t["to-delete"] = Vertex{0,0}

	// Delete an element delete(m, key)
	delete(t, "to-delete")

	fmt.Println(t)
	// Retrieve an element
	if elem, ok := m["to-delete"]; ok{
		fmt.Println(elem)
	}else{
		fmt.Println("m['to-delete'] dont exists")
	}
}


// cast recieve three arguments, x, y (ints) and a function which is used infunction
func cast(x int, y int, funcVar func(int, int)int)string{
	return strconv.Itoa(funcVar(x,y))
}
// FunctionValues functions are values too. They can be passed around just like other values
func FunctionValues(){
	// Function values can be used as function arguments and return values
	

	sum := func(x, y int) int{
		return x + y
	}

	fmt.Println("call function as sum(1,2)")
	fmt.Println(sum(1,2))
	fmt.Println("send function as argument")
	fmt.Println(cast(3,4,sum))
}

// This function serves as the closure
func adder() func(int) int{
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}
// Closures are functions values that references variables from outside its body. The function may access and assign to the referenced variables 
func Closures(){
	pos := adder()
	for x:=0; x <= 10; x++{
		fmt.Println(pos(x))
	}
}

func fibonacciClosure() func() int{
	x, y := 0, 1
	return func() int{
		x, y = y, x+y
		return x
	}
}

// Fi function calls a closure that creates the fibonacci series
func Fi(){
	f := fibonacciClosure()
	for i:= 0; i< 10; i++{
		fmt.Println(f())
	}
}