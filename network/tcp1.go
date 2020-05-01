package network

import (
	"flag"
)

// SetupTCP1 is used to setup the TCP server
// specifying wherther to start up server or client
func SetupTCP1() {
	what := flag.String("mode", "server", "Specify what you want to run: server or Client")
	flag.Parse()
	if *what == "server" {
		SetupTCPServer()
	} else {
		SetupTCPClient()
	}
	// switch what {
	// case "client":
	// 	SetupClientServer()
	// 	breakt
	// default:
	// 	SetupTcpServer()
	// 	break
	// }
}
