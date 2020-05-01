package presuffix

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"math"
	"os"
)

const n int = 10000

// Compress is used to check how compress works
func Compress() {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	fmt.Println(string(buff))
	b := bytes.NewReader(buff)

	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	io.Copy(os.Stdout, r)
	// Output: hello, world
	r.Close()
}

// GetPi is a function to get pi
func GetPi() {
	// runtime.GOMAXPROCS(1) // in goroutine_select2.go
	ch := make(chan float64)
	go generateN(ch)
	getData(ch)
	// var val float64
	// for i := 0; i < n; i++ {
	// 	val += <-ch
	// }
	// fmt.Println(val)
	// time.Sleep(9e9)
	// var val float64
	// for {
	// 	val += <-ch
	// }
	// fmt.Println(val)
}

func formula(k float64) float64 {
	top := math.Pow(-1, k)
	bottom := (2*k + 1)
	return (top / bottom)
}
func generateN(ch chan float64) {
	for i := 0; i < n; i++ {
		go func(i int) {
			val := 4 * formula(float64(i))
			// fmt.Printf("value: %v\t index: %v \n", val, i)
			ch <- val
		}(i)
	}
}

func getData(ch chan float64) {
	var val float64
	for i := 0; i < n; i++ {
		val += <-ch
	}
	fmt.Println(val)
}
