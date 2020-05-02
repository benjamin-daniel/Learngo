package clip

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
)

// CopyOverTime prints the clipboard to the console
// every 1 second
func CopyOverTime() {
	for {
		s, err := clipboard.ReadAll()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("%s\n", s)
		time.Sleep(1000 * time.Millisecond)
	}
}
