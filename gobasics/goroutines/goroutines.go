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

