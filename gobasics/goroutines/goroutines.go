package goroutines

import(
	"fmt"
	"time"
)

func say(s string){
	for i := 0; i < 5; i++{
		fmt.Printf("%v <- \n", s)

		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("\t <- %v\n", s)
	}
}
// BasicGoroutine is a lightweight thread managed by the go runtime
func BasicGoroutine(){
	// Starts a new go routine
	go say("World")
	//go say("test")
	say("Hello")
}


func sum(s []int, c chan int){
	sum := 0
	for _, val := range s{
		sum += val
	}
	// Send value sum to channel
	c <- sum  // The data flows in the direction of the arrow
}
// BasicChannels channels are a type of conduct through which you can send and receive values with the channel operator <-
func BasicChannels(){

	s := []int{3,4,8,10,-3,-100}
	// Like maps and slices, channels must be created before use "ch := make(chan int)"
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c //receive from channel and assign 
	// Once both goroutines have completed their computation, it calculates the final result

	// By default, sends and receives block until the other side is ready. 
	// This allows goroutines to synchronize without explicit locks or condition variables
	
	fmt.Println(x, y, x+y)
}
