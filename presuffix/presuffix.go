package presuffix

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/wcharczuk/go-chart"
)

var (
	ii uint64
	dd float64
	ss string
)

func toBase2AndReverse(num float64) (result float64) {
	n := int64(num)
	base2 := reverse(strconv.FormatInt(n, 2))
	resultInt64, _ := strconv.ParseInt(base2, 2, 64)
	return float64(resultInt64)
}

func generatePrimes(num int) (primes []float64) {
	var x, y, n int
	nsqrt := math.Sqrt(float64(num))
	isPrime := make([]bool, num)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= num && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= num && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= num && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < num; y += n * n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	primes = make([]float64, 0, 1270606)
	for x = 0; x < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, float64(x))
		}
	}
	return primes
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
func makeRange(min, max int) []float64 {
	a := make([]float64, max-min+1)

	for i := range a {
		a[i] = float64(min) + float64(i)
	}
	return a
}
func sieveOfEratosthenes(N int) (primes []float64) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, float64(i))
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

// DrawBinaryPrimes is used to test my String prowress
func DrawBinaryPrimes() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(toBase2AndReverse(i))

	// }
	yAxis := make([]float64, 0, 1270606)
	xAxis := sieveOfEratosthenes(100000)
	// fmt.Println(xAxis)
	for _, val := range xAxis {
		yAxis = append(yAxis, toBase2AndReverse(val))
	}
	// fmt.Println(yAxis)
	// fmt.Println(xAxis)
	// // fmt.Printf("primes: %v", )
	// fmt.Println(yAxis)
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xAxis,
				YValues: yAxis,
			},
		},
	}

	// buffer := bytes.NewBuffer([]byte{})
	buffer, _ := os.Create("DrawBinaryPrimes.png")

	err := graph.Render(chart.PNG, buffer)
	// fmt.Println(buffer)
	// f, err := os.Create("image.png")

	// png.Encode(f, buffer)

	fmt.Println(err)

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// Declare second integer, double, and String variables.
	// var (
	// 	i uint64  = 232
	// 	d float64 = 232
	// 	s string  = "Hey"
	// )

	// Read and save an integer, double, and String to your variables.
	// Read and save an integer, double, and String to your variables.

}

// UseOSStin is used to test my os.stdin
func UseOSStin() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Interger: ")
	scanner.Scan()
	iii := scanner.Text()
	ii, _ = strconv.ParseUint(iii, 10, 64)
	fmt.Print("Enter Number: ")
	scanner.Scan()
	ddd := scanner.Text()
	dd, _ = strconv.ParseFloat(ddd, 64)
	fmt.Print("Enter String: ")
	scanner.Scan()
	ss = scanner.Text()
	// Print the sum of both integer variables on a new line.
	fmt.Printf("%v \n", ii)
	fmt.Printf("%v \n", dd)
	fmt.Printf("%v \n", ss)
}
