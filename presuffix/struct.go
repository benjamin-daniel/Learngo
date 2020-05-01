package presuffix

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

type address struct {
	name    string
	address string
}

type vCard struct {
	name      string
	addresses []*address
	photo     string
}

// TestStruct is a function for initilizing structs
func TestStruct() {
	dan := user{name: "Benjamin Daniel", age: 10}
	micheal := new(user)
	micheal.age = 12
	micheal.name = "Micheal"
	fav := &user{"Favour", 16}
	fmt.Println(micheal)
	fmt.Println(fav)
	fmt.Println(dan)
	fadeyi := &address{
		address: "Fadeyi",
	}
	fmt.Println(fadeyi)
	DanisVCard := vCard{
		name:      "Dani",
		photo:     "profile.jpg",
		addresses: []*address{fadeyi},
	}
	fmt.Println(DanisVCard)
}
