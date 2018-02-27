// Package flowControl covers for, if, switch and defer Structtures
package flowControl

import (
	"time"
	"fmt"
	"os"
	"log"
	"runtime"
)

// ForStruct covers the simple Structture and the reduced one
func ForStruct (limit int){
	// Simple for (unlike many languages, Go dosn't require the parentesis)
	for i:=0; i<limit; i++{
		fmt.Println("for i:=0; i<5; i++{} // ", i)
	}

	// ReducedFor only takes the condition expression without an init nor post statement but both must exist elsewhere	
	// The ReducedFor works as a while
	i := 1
	for i < limit {
		fmt.Println("for i < limit {} // ", i)
		i+=i
	}
}

// IfStruct covers the simple Structture
func IfStruct (x int){
	// SimpleIf syntax like a for loop // Since error cases tend to end in return statements, the resulting code needs no else statements.
	if y := x; y > 5{
		fmt.Printf("%v is greater than 5\n", y)
	}else{
		fmt.Printf("%v is lesser than 5\n", y)
	}

	// IfErrorCatching basic error catching Structture
	wd, err := os.Getwd()
		// Log Working directory
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
		//Code using f and d
	fmt.Println(f, " and ", d)
}

// SwitchStruct covers the simple Struct, the evaluation order and the switch with no condition
func SwitchStruct(){
	// SimpleSwitch a switch statement is a shorter way to write a secuence of if-else statements
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%v \n", os)
	}

	// SwitchEvaluationOrder shows the order top-bottom of the switch evaluations
	fmt.Println("When's saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	// SwitchNoCondition a switch without condition is the same as a switch-true
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good night")
	}
}

// DeferStruct defers the execution of a function until the surrounding functions return (LIFO)
func DeferStruct(){

	// StackedDefer defer functions are pushed onto a stack where Last-In-First-Out
	for i:= 0 ; i < 10; i++{
		defer fmt.Println(i)
	}
	
	defer fmt.Printf("World\n")
	defer fmt.Print(" ")
	fmt.Print("Hello")
}
