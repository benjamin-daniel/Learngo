package presuffix

import "fmt"

func generate(ch chan int) {
	for i := int(1e18); ; i++ {
		ch <- i // Send i to channel ch.
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable i from in.
		if i%prime != 0 {
			out <- i // Send i to channel out.
		}
	}
}

// Sieve uses goroutines to make a algo for generating and sieving prime numbers
func Sieve() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
