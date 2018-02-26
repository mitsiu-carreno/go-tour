// Package flowControl covers for, if, switch and defer structures
package flowControl

import (
	"time"
	"fmt"
	"os"
	"log"
	"runtime"
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
	//Code using f and d
	fmt.Println(f, " and ", d)
}

// SimpleSwitch a switch statement is a shorter way to write a secuence of if-else statements
func SimpleSwitch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%v \n", os)
	}
}

// SwitchEvaluationOrder shows the order top-bottom of the switch evaluations
func SwitchEvaluationOrder(){
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

}

// SwitchNoCondition a switch without condition is the same as a switch-true
func SwitchNoCondition(){
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

// Defer defers the execution of a function until the surrounding functions return (LIFO)
func Defer(){
	defer fmt.Printf("World\n")
	defer fmt.Print(" ")
	fmt.Print("Hello")
}

// StackedDefer defer functions are pushed onto a stack where Last-In-First-Out
func StackedDefer()  {
	for i:= 0 ; i < 10; i++{
		defer fmt.Println(i)
	}
}