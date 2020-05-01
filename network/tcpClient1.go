package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// SetupTCPClient is used to set up a client tcp server to push things to the main server
func SetupTCPClient() {
	fmt.Println("opening  connection")
	// create the listerner
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error Dialling", err.Error())
		return // terminate program
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name? ")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	_, err = conn.Write([]byte(trimmedClient + " joined"))
	connClose := func() {
		_, err = conn.Write([]byte(trimmedClient + " left"))
		conn.Close()
		return
	}

	defer connClose()
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		// fmt.Printf("input:--%s--",input)
		// fmt.Printf("trimmedInput:--%s--",trimmedInput)
		if trimmedInput == "!Q" {
			connClose()
		}
		// send the byte of data to the server
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))

	}

}

// func main() {
// 	SetupTCPClient()
// }
