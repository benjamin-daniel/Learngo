package network

import (
	"fmt"
	"net"
)

// SetupTCPServer is used to set up a tcp server on port 5000
func SetupTCPServer() {
	fmt.Println("Starting the server")
	// create the listerner
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // terminate program
	}
	// listen and accept connections from clients
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting", err.Error())
			return // terminate program
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // terminate program
		}
		fmt.Printf("Received data: %v\n", string(buf))
	}

}
