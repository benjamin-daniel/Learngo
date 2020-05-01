package presuffix

import "fmt"

// TestMapRef is used to test Map references
func TestMapRef() {
	var mapAssigned map[string]int
	mapLit := map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3 // this changes mapLit value
	// fmt.Println(mapLit)
	// fmt.Println(mapAssigned)
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}

// FuncAssigned is used to see that functions can be assigned as values
func FuncAssigned() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
	fmt.Println(mf[1]())

}