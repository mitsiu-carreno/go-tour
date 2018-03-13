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
// FunctionReturnChannel channels as a handle on a service 
func FunctionReturnChannel(){
	// Our boring function returns a channel that let us communicate with the boring service it provides 
	c := boringFunc("Boring func")
	for i:= 0; i<5; i++{
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You are a boring func, I'm leaving")

	joe := boringFunc("Joe")
	ann := boringFunc("Ann")
	for i := 0; i<5; i++{
		// They sycronize, if joe is ready to send but ann isn't, joe is blocked until ann send
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}
	fmt.Println("You are both boring services, I'm leavnig")
}