package presuffix

import (
	"fmt"
	"time"
)

// TestTimer is to test the goroutines produced by the time package
func TestTimer() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}

}
