// Package flowControl covers for, if, switch and defer structures
package flowControl

import (
	"fmt"
	"os"
	"log"
)

// SimpleFor unlike many languages, Go dosn't require the parentesis
func SimpleFor (){
	for i:=0; i<5; i++{
		fmt.Println("for i:=0; i<5; i++{} // ", i)
	}
}

// ReducedFor only takes the condition expression without an init nor post statement but both must exist elsewhere
func ReducedFor(limit int){
	i := 1
	for i < limit {
		fmt.Println("for i < limit {} // ", i)
		i+=i
	}
}

// ForAsWhile is as simple as the ReducedFor
func ForAsWhile(limit int){
	fmt.Println("use a ReducedFor as a while in go")
	ReducedFor(limit)
}

// SimpleIf syntax like a for loop // Since error cases tend to end in return statements, the resulting code needs no else statements.
func SimpleIf(x int){
	if y := x; y > 5{
		fmt.Printf("%v is greater than 5\n", y)
	}else{
		fmt.Printf("%v is lesser than 5\n", y)
	}
}

// IfErrorCatching basic error catching structure
func IfErrorCatching(){
	wd, err := os.Getwd()
	log.Println(wd)
	
	f, err := os.Open("gobasics/flowControl/test.txt")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	d, err := f.Stat()
	if err != nil {
		f.Close()
		log.Println(err)
	}

	fmt.Println(f, d)
}