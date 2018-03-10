// Package concurrencyPatterns is a series of examples extracted from (https://www.youtube.com/watch?v=f6kdp27TYZs)
package concurrencyPatterns

import (
	"fmt"
	"time"
	"math/rand"
)

func boring (msg string){
	for i:=0; i<10; i++{
		fmt.Println(msg, i)
		time.Sleep((time.Duration(rand.Intn(1e3))) * time.Millisecond)
	}
}

// Concurrency is the composition of independently executing computations
func Concurrency(){
	go boring("boring!")
	fmt.Println("I'm listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving")
}