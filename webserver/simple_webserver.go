package webserver

import (
	"io"
	"net/http"
)

const form = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>First Golang Project</title>
  </head>
  <body>
    <form action="#" method="post" name="bar">
      <input type="text" name="in" />
      <input type="submit" value="Submit" />
    </form>
  </body>
</html>
`

// SimpleServer is used to handle a simple get request
func SimpleServer(w http.ResponseWriter, request *http.Request) {
	// w.Header().Set("Content-Type", "text/txt") // would make the browser interpret this as a txt reponse

	io.WriteString(w, "<h1>hello, world</h1>")
}

// FormServer is for handling the form issue
func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data*/
		//request.ParseForm();
		io.WriteString(w, request.FormValue("in"))
	}
}

// WorkWithForm is the web server function
func WorkWithForm() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
