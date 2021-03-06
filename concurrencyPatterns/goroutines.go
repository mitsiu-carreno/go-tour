// Package concurrencyPatterns is a series of examples extracted from (https://www.youtube.com/watch?v=f6kdp27TYZs)
package concurrencyPatterns

// DONT COMMUNICATE BY SHARING MEMORY, SHARE MEMORY BY COMMUNICATING

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
	fmt.Println("I'm listening for two seconds")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving")
}

func boringChannel(s string, c chan string){
	for i:=0; ; i++{
		c<- fmt.Sprintf("%s %d", s, i)
		time.Sleep((time.Duration(rand.Intn(1e3))) * time.Millisecond)
	}
}
// Channel provides a connection between two goroutines, allowing them to communicate, (it syncronize the routines)
func Channel(){
	c := make(chan string)
	go boringChannel("boring", c)
	for i:=0; i<5; i++{
		fmt.Printf("You say: %q\n", <-c);
	}
	fmt.Println("You're boring i'm leaving")
}