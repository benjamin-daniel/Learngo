package webserver

import (
	"fmt"
	"log"
	"net/http"
)

// HelloServer is a handler
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside helloserver handler")
	fmt.Println("req: ", req)
	fmt.Println(req.URL.Path)
	fmt.Fprint(w, "Hello,"+req.URL.Path[1:])
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):] //get everything after the /hello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

// StartHelloWorldServer starts a server on port 6767
func StartHelloWorldServer() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:6767", nil)
	if err != nil {
		log.Fatal("Listen and Server: ", err.Error())
	}
}
