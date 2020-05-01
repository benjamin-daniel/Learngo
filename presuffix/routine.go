package presuffix

import (
	"fmt"
	"time"
)

// RunGoRoutine is used to run my codes testing RunGoRoutine
func RunGoRoutine() {
	fmt.Println("In main()")
	now := time.Now()
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	fmt.Println(now)
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}
func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds fmt.Println("End of longWait()")
}
func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds fmt.Println("End of shortWait()")
}

// ChannelBlock is for showing a channel block
func ChannelBlock() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	time.Sleep(1e9)
}

func pump(ch1 chan int) {
	for i := 0; ; i++ {
		ch1 <- i
	}
}
func suck(ch1 chan int) {
	for {
		fmt.Println(<-ch1)
	}
}
