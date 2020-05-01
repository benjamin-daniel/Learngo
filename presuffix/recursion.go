package presuffix

import "fmt"

// Rec is a function used to test recursion
func Rec() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// 18 is odd: is false
}
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(revSign(nr) - 1)
}
func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(revSign(nr) - 1)
}
func revSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
