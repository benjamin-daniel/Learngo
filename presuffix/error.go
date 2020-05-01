package presuffix

import "fmt"

// CatchError is used to catch a real error
func CatchError(num int) {

	defer func() {
		if r := recover(); r != nil {
			err, _ := r.(error)
			fmt.Println(err)
			// if !ok {
			// 	err = fmt.Errorf("Pkg Error: %v", err)
			// }
		}
	}()

	val := 1 / num
	fmt.Println(val)
	return
}
