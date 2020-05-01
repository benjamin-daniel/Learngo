package presuffix

import (
	"fmt"
)

// func addf(x int) func {
// 	return func( y int)  int{
// 		return x + y
// 	}
// }
type funk func(int) int

func addf(x int) (f funk) {
	fmt.Println("First")
	return func(y int) int {
		fmt.Println("Second")
		return x + y
	}
}

func fUnk() {
	// add5 := addf(5)
	// fmt.Println(add5(6))
	x := Min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	// fmt.Println(arr.(type))
	x = Min(arr...)
	fmt.Printf("The minimum in the array arr is: %d\n", x)
}

// Min is used to find the minimum of two or more numbers
func Min(a ...int) (min int) {
	if len(a) == 0 {
		return 0
	}
	min = a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
