package presuffix

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"unicode/utf8"
)

// Arr is for testing array things
func Arr() {
	s := make([]byte, 5)
	m := s
	s = s[2:4]
	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(cap(s))
}

// inspectSlice exposes the slice header for review.
// Parameter: again, there is no value in side the []string so we want a slice.
// Range over a slice, just like we did with array.
// While len tells us the length, cap tells us the capacity
// In the output, we can see the addresses are aligning as expected.
// func inspectSliceInt(slice []int) {
// 	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
// 	for i := range slice {
// 		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
// 	}
// }
func inspectSliceString(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}

// CrazySlice is a slice problem that has been giving me issues understanding
func CrazySlice() {

	// ------------------
	// Length vs Capacity
	// ------------------

	// Length is the number of elements from this pointer position we have access to (read and write).
	// Capacity is the total number of elements from this pointer position that exist in the
	// backing array.
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	var ars = ar[4:8]
	ars = append(ars, 4, 3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(a)) // 2
	fmt.Println(cap(a)) // 5
	fmt.Println(ars)
	fmt.Println(len(ars))
	fmt.Println(cap(ars))
	fmt.Println(ar[:2])
	fmt.Println(ar[3:])
	fmt.Println(append(ar[:2], ar[3:]...))
	// inspectSliceInt(ars)
}

// Deep is a function for testing deep theory
func Deep() {
	// -----------------------------
	// Contiguous memory allocations
	// -----------------------------

	// Declare an array of 6 strings initialized with values.
	six := [6]string{"Annie", "Betty", "Charley", "Doug", "Edward", "Hoanh"}

	// Iterate over the array displaying the value and address of each element.
	// By looking at the output of this Printf function, we can see that this array is truly a
	// contiguous block of memory. We know a string is 2 word and depending on computer
	// architecture, it will have x byte. The distance between two consecutive IndexAddr is exactly
	// x byte.
	// v is its own variable on the stack and it has the same address every single time.
	fmt.Printf("\n=> Contiguous memory allocations\n")
	for i, v := range six {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &six[i])
	}

	// --------------
	// Reference type
	// --------------

	// Create a slice with a length of 5 elements and a capacity of 8.
	// make allows us to adjust the capacity directly on construction of this initialization.
	// What we end up having now is a 3 word data structure where the first word points to an array
	// of 8 elements, length is 5 and capacity is 8.
	//  -----
	// |  *  | --> | nil | nil | nil | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |  0  |  0  |  0  |
	// |  5  |
	//  -----
	// |  8  |
	//  -----
	// It means that I can read and write to the first 5 elements and I have 3 elements of capacity
	// that I can leverage later.
	slice2 := make([]string, 5, 8)
	slice2[0] = "Apple"
	slice2[1] = "Orange"
	slice2[2] = "Banana"
	slice2[3] = "Grape"
	slice2[4] = "Plum"

	fmt.Printf("\n=> Length vs Capacity\n")
	inspectSliceString(slice2)

	// --------------------------------------------------------
	// Idea of appending: making slice a dynamic data structure
	// --------------------------------------------------------
	fmt.Printf("\n=> Idea of appending\n")
	// Declare a nil slice of strings, set to its zero value.
	// 3 word data structure: first one points to nil, second and last are zero.
	var data []string

	// What if I do data := string{}? Is it the same?
	// No because data in this case is not set to its zero value.
	// This is why we always use var for zero value because not every type when we create an empty
	// literal we have its zero value in return.
	// What actually happen here is that we have a slice but it has a pointer (as opposed to nil).
	// This is consider an empty slice, not a nil slice.
	// There is a semantic between a nil slice and an empty slice. Any reference type that set to
	// its zero value can be considered nil. If we pass a nil slice to a marshal function, we get
	// back a string that said null but when we pass an empty slice, we get an empty JSON document.
	// But where does that pointer point to? It is an empty struct, which we will review later.

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~100k strings to the slice.
	for record := 1; record <= 102400; record++ {
		// Use the built-in function append to add to the slice.
		// It allows us to add value to a slice, making the data structure dynamic, yet still
		// allows us to use that contiguous block of memory that gives us the predictable access
		// pattern from mechanical sympathy.
		// The append call is working with value semantic. We are not sharing this slice but
		// appending to it and returning a new copy of it. The slice gets to stay on the stack, not
		// heap.
		data = append(data, fmt.Sprintf("Rec: %d", record))

		// Every time append runs, it checks the length and capacity.
		// If it is the same, it means that we have no room. append creates a new backing array,
		// double its size, copy the old value back in and append the new value. It mutates its copy
		// on its stack frame and return us a copy. We replace our slice with the new copy.
		// If it is not the same, it means that we have extra elements of capacity we can use. Now we
		// can bring these extra capacity into the length and no copy is being made. This is very
		// efficient.

		// Looking at the last column in the output, when the backing array is 1000 elements or
		// less, it doubles the size of the backing array for growth. Once we pass 1000 elements,
		// growth rate moves to 25%.

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {
			// Calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for capacity.
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
		}
	}
}

// Sum is used to calculate the sum of an array
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		//  not necessary!
		sum += v
	}
	return
}

var digitRegexp = regexp.MustCompile("[0-9]+")

// TestMany is a function to test differnent slice methods
func TestMany() {
	google := "Google"
	googleSlice := []byte(google)
	sf := googleSlice[0]
	sl := googleSlice[len(google)-1]
	googleSlice[0], googleSlice[len(google)-1] = sl, sf
	google = string(googleSlice)
	fmt.Println(google)

	fmt.Println(findDigits("slice.txt"))
}

func findDigits(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// Trouble shows the bad side of slice and reference
func Trouble() {

	// -------------------
	// Slice and reference
	// -------------------

	// Declare a slice of integers with 7 values.
	x := make([]int, 7)

	// Random starting counters.
	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}

	// Set a pointer to the second element of the slice.
	twohundred := &x[1]

	// Append a new value to the slice. This line of code raises a red flag.
	// We have x is a slice with length 7, capacity 7. Since the length and capacity is the same,
	// append doubles its size then copy values over. x nows points to diffrent memeory block and
	// has a length of 8, capacity of 14.
	x = append(x, 800)

	// When we change the value of the second element of the slice, twohundred is not gonna change
	// because it points to the old slice. Everytime we read it, we will get the wrong value.
	x[1]++

	// By printing out the output, we can see that we are in trouble.
	fmt.Printf("\n=> Slice and reference\n")
	fmt.Println("twohundred:", *twohundred, "x[1]:", x[1])
}

// Utf8 is used to test bytes
func Utf8() {
	// -----
	// UTF-8
	// -----
	fmt.Printf("\n=> UTF-8\n")

	// Everything in Go is based on UTF-8 character sets.
	// If we use different encoding scheme, we might have a problem.

	// Declare a string with both Chinese and English characters.
	// For each Chinese character, we need 3 byte for each one.
	// The UTF-8 is built on 3 layers: bytes, code point and character. From Go perspective, string
	// are just bytes. That is what we are storing.
	// In our example, the first 3 bytes represents a single code point that represents that single
	// character. We can have anywhere from 1 to 4 bytes representing a code point (a code point is
	// a 32 bit value) and anywhere from 1 to multiple code points can actually represent a
	// character. To keep it simple, we only have 3 bytes representing 1 code point representing 1
	// character. So we can read s as 3 bytes, 3 bytes, 1 byte, 1 byte,... (since there are only 2
	// Chinese characters in the first place, the rest are English)
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune -> maximum number of bytes we need to
	// represent any code point is 4.
	// Rune is its own type. It is an alias for int32 type. Similar to type byte we are using, it
	// is just an alias for uint8.
	var buf [utf8.UTFMax]byte

	// When we are ranging over a string, are we doing it byte by byte or code point by code point or
	// character by character?
	// The answer is code point by code point.
	// On the first iteration, i is 0. On the next one, i is 3 because we are moving to the next
	// code point. Then i is 6.
	for i, r := range s {
		// Capture the number of bytes for this rune/code point.
		rl := utf8.RuneLen(r)

		// Calculate the slice offset for the bytes associated with this rune.
		si := i + rl

		// Copy rune from the string to our buffer.
		// We want to go through every code point and copy them into our array buf, and display
		// them on the screen.
		// "Every array is just a slice waiting to happen." - Go saying
		// We are using the slicing syntax, creating our slice header where buf becomes the backing
		// array. All of them are on the stack. There is no allocation here.
		copy(buf[:], s[i:si])

		// Display the details.
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}
