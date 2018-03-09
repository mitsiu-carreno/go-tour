// concurrencyPatterns 
package concurrencyPatterns

import (
	"fmt"
	"time"
	_"math/rand"
)

func boring (msg string){
	for i:=0; i<10; i++{
		fmt.Println(msg, i)
		time.Sleep(2 * time.Millisecond)
	}
}

func Index(){
	go boring("boring!")
	fmt.Println("I'm listening")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("You're boring; I'm leaving")
}