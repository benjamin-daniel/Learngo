package reader

import (
	"log"
	"os"
)

// Reader is for demonstrate how io.Reader works
func Reader() {
	file, err := os.OpenFile("reader_interface/sloth.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("error opening sloth.txt: %v", err)
	}
	defer file.Close()

	// Make a byte slice that's big enough to store a few words of the message
	// we're reading
	bytesRead := make([]byte, 33)
	// 33 isn't the length of the txt file so it only stores 33byte of data

	// Now read some data, passing in our byte slice
	n, err := file.Read(bytesRead)
	if err != nil {
		log.Fatalf("error reading from sloth.txt: %v", err)
	}

	// Take a look at the bytes we copied into the byte slice
	log.Printf("We read \"%s\" into bytesRead (%d bytes)",
		string(bytesRead), n)
	n, err = file.Read(bytesRead)
	if err != nil {
		log.Fatalf("error reading from sloth.txt: %v", err)
	}

	log.Printf("We read \"%s\" into bytesRead (%d bytes)",
		string(bytesRead), n)
}
