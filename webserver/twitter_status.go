package webserver

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* these structs will house the unmarshalled response.
         they should be hierarchically shaped like the XML
but can omit irrelevant data. */

// Status is one of the things I don't understand in this program
type Status struct {
	Text string
}

// User is a struct of user
type User struct {
	XMLName xml.Name
	Status  Status
}

// FetchUser is used to fetch a twitter user
// twitter has changed things
// this doesn't work anymore
func FetchUser() {
	// perform an HTTP request for the twitter status of user: Googland
	response, err := http.Get("http://twitter.com/users/Googland.xml")
	if err != nil {
		log.Fatalln("Fetch Error: ", err.Error())
	}
	defer response.Body.Close()
	// initialize the structure of the XML response
	user := User{xml.Name{"", "user"}, Status{""}}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	fmt.Println(string(body))

	// unmarshal the XML into our structures
	xml.Unmarshal(body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}
