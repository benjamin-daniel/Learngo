package presuffix

import (
	"fmt"
	"time"
)

// IntensePlay is a function to test out what I'm watching in
func IntensePlay() {
	// var ch1 chan int
	ch1 := make(chan int)
	// ch1 <- 1
	go worker(ch1)
	ch1 <- 1
	fmt.Printf("c1 is %T %v \n", ch1, ch1)
}

func worker(ch chan int) {
	fmt.Println("Worker")
	<-ch
}

// BufferedCh is a func to test buffered channels
func BufferedCh() {
	ch := make(chan int, 3)
	go worker(ch)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
}

func getValuesFromBufferedCh(ch chan int) {
	fmt.Println(<-ch)
}

// ------------------------------------------
// Unbuffered channel: Signaling without data
// ------------------------------------------

// SignalClose shows how to close a channel to signal an event.
func SignalClose() {
	// We are making a channel using an empty struct. This is a signal without data.
	ch := make(chan struct{})

	// We are gonna launch a Goroutine to do some work. Suppose that it's gonna take 100
	// millisecond. Then, it wants to signal another Goroutine that it's done.
	// It's gonna close the channel to report that it's done without the need of data.

	// When we create a channel, buffered or unbuffered, that channel can be in 2 different states.
	// All channels start out in open state so we can send and receive data. When we change the
	// state to be closed, it cannot be opened. We also cannot close the channel twice because that
	// is an integrity issue. We cannot signal twice without data twice.

	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("signal event")
		close(ch)
	}()

	// When the channel is closed, the receive will immediately return.
	// When we receive on a channel that is open, we cannot return until we receive the data
	// signal. But if we receive on a channel that is closed, we are able to receive the signal
	// without data. We know that event is occurred. Every receive on that channel will immediately
	// return.
	println("hey")
	<-ch

	fmt.Println("event received")
}
