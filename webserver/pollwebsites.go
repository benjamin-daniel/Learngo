package webserver

import (
	"fmt"
	"log"
	"net/http"
)

var urls = []string{
	"http://google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

// PollUrls is used to check urls and see how they react
func PollUrls() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			log.Println("Error: ", url, err)
		} else {
			fmt.Println(url+" : ", resp.Status)
		}
	}
}
