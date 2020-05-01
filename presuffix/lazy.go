package presuffix

import (
	"fmt"
)

// Any is an interface to allow me accept any type in my functions
type Any interface{}

// EvalFunc is a function that accepts anyvalue and returns any value
type EvalFunc func(Any) (Any, Any)

// Lazy is a function to test and use interfaces and channels
// this is a request-response mechanism
func Lazy() {
	// // evenFunc is a function that that returns the number put in and the next even number
	// evenFunc := func(state Any) (Any, Any) {
	// 	// converts the interface to integer (type convertion)
	// 	os := state.(int)
	// 	// fmt.Println("os:", os)
	// 	// this computes the next value
	// 	ns := os + 2
	// 	// this returns the current value then the next value
	// 	return os, ns
	// }

	// // even is a function
	// even := BuildLazyIntEvaluator(evenFunc, 0)
	// even()
	fibb := func(n Any) (Any, Any) {
		// intn := n.(int)
		// power2 := int(math.Pow(float64(intn), 2))
		// return intn, power2
		// n = n.(uint64)
		intn := n.(int)
		n = uint64(intn)

		// current holds the nth Fibonacci number.
		current := 0
		next := current

		// // return 0 1 or for base cases (ideally)
		// if n == 0 || n == 1 {
		// 	return n
		// }
		switch n {
		case 0:
			return 0, 1
		case 1:
			return 1, 1
		}

		f0 := 0
		f1 := 1

		for i := 0; i < intn-1; i++ {
			current = f0 + f1
			f0 = f1
			f1 = current
			if i == intn-2 {
				// fmt.Printf("Mutated i: %v, cond: %v \n", i, n-1)
				next = f0 + f1
			}
		}

		return current, next
	}
	fibb(0)
	fibo := BuildLazyIntEvaluator(fib, 0)
	for i := 0; i < 10; i++ {
		// now, later := fib(i)
		// fmt.Printf("Now: %v \t Later: %v\n", now, later)
		// fmt.Printf("%vth even: %v\n", i, even())
		fmt.Printf("%vth fibo: %v\n", i, fibo())
	}
}
func fib(n Any) (Any, Any) {
	// current holds the nth Fibonacci number.
	intn := n.(int)
	n = uint64(intn)

	// current holds the nth Fibonacci number.
	current := 0

	// // return 0 1 or for base cases (ideally)
	// if n == 0 || n == 1 {
	// 	return n
	// }
	switch n {
	case 0:
		return 0, 1
	case 1:
		return 1, 1
	}
	// keep track of 2 previous numbers at each step to build the Fibonacci
	// series from the bottom up
	f0 := 0
	f1 := 1

	// to get nth Fibonacci, do n-1 iterations.
	for i := 0; i < intn-1; i++ {
		current = f0 + f1
		f0 = f1
		f1 = current
	}

	return current, intn + 1
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {

	// retValChan is basically return value channel
	// this is the channel we would use for communication
	retValChan := make(chan Any)
	// loopFunc is a function used to
	loopFunc := func() {
		// actState(Active State) is a copy of the initial state
		var actState Any = initState
		// retVal(return value) is what holds the evaluated function
		var retVal Any
		for {
			// we then set retVal to the value of the argument function
			// actState is then set to the next value
			retVal, actState = evalFunc(actState)
			// fmt.Printf("Now: %v \t Later: %v\n", retVal, actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}
func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}
