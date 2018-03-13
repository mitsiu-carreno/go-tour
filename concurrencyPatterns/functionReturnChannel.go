package concurrencyPatterns

import (
	"fmt"
	"time"
	"math/rand"
)

func boringFunc(msg string) <-chan string{	// Returns receive-only channel of string
	c := make(chan string)
	go func (){ // We launch the goroutine from inside the function
		for i:=0; ; i++{
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3))* time.Millisecond)
		}
	}()
	return c // Return the channel to the caller

}
// FunctionReturnChannel channels are first-class values, just like strings or integers
func FunctionReturnChannel(){
	c := boringFunc("Boring func")
	for i:= 0; i<5; i++{
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are a boring func, I'm leaving")
}