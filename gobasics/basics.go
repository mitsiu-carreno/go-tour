// Package gobasics contains basics conepts from the go guide
package gobasics

import "fmt"

// NakedReturn name the return variables so you just call an 'return' statement
func NakedReturn(x, y string, z int) (a, b string, c int) {
	a, b, c = x, y, z
	return
}

// TypeConversions follow the expression T(v) (with value v and type T)
func TypeConversions(x int) (float64, uint){
	var f = float64(x)
	var u = uint(f)
	return f, u
}

//TypeInterface if no type is declared explicitly, the variable type is inferred from the value on the right hand side
func TypeInterface (){
	v := 45
	var i float64 = 45
	j := i
	fmt.Printf("v := 45 (v type %T) // var i float64 = 45; j := i (j type %T) \n", v, j)
}