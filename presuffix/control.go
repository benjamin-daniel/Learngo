package presuffix

import (
	"fmt"
	"strings"
)

// Control is the a function I use to run main control file function
func Control() {
	// fallThrough()
	// forCharacter(2)
	fmt.Println(multiply(2, 3, 4, 5))
}

func forCharacter(n int) (err string) {
	if n < 1 || n > 25 {
		err = "Can't pass more than 25 or less than 1"
		return
	}
	for ix := 1; ix <= n; ix++ {
		fmt.Println(strings.Repeat("G", ix))
	}
	return
}

func multiply(a ...int) (n int) {
	n = 1
	for i := 0; i < len(a); i++ {
		n = n * a[i]
	}
	return
}
func fallThrough() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
