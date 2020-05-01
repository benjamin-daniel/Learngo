package presuffix

import (
	"fmt"
	"math"
)

type shaper interface {
	Area()
}

type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (s *square) Area() {
	area := math.Pow(s.length, 2)
	fmt.Println("The value of areat is ", area)
}

func (s *circle) Area() {
	area := math.Pow(s.radius, 2) * math.Pi
	fmt.Println("The value of areat is ", area)
}

// Trying to do this from error package
// type timeout interface {
// 	Timeout() bool
// }

// // PathError records an error and the operation and file path that caused it.
// type PathError struct {
// 	Op   string
// 	Path string
// 	Err  error
// }

// func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

// func (e *PathError) Unwrap() error { return e.Err }

// // Timeout reports whether this error represents a timeout.
// func (e *PathError) Timeout() bool {
// 	t, ok := e.Err.(timeout)
// 	return ok && t.Timeout()
// }

// InterfaceFreaky is used to make some test some freaky things in interface
func InterfaceFreaky() {
	firstCircle := circle{radius: 4}
	// firstCircle = shaper{firstCircle}
	fmt.Println(firstCircle)
	// t, ok := firstCircle.(shaper)
	// fmt.Println(t, ok)
}
