package main

import (
	reader "github.com/benjamin-daniel/hello/reader_interface"
)

func main() {
	// presuffix.Control()
	// presuffix.Arr()
	// presuffix.CrazySlice()
	// presuffix.Deep()
	// presuffix.TestMany()
	// presuffix.Utf8()
	// presuffix.TestMapRef()
	// presuffix.FuncAssigned()
	// presuffix.TestStruct()
	// presuffix.ReadFileFromBuffer()
	// presuffix.CatchError(0)
	// presuffix.RunGoRoutine()
	// presuffix.ChannelBlock()
	// presuffix.Sieve()
	// presuffix.GetPi()
	// presuffix.CatchError(0)
	// presuffix.InterfaceFreaky()
	// presuffix.Compress()
	// presuffix.TestTimer()
	// presuffix.GoRoutine1()
	// presuffix.GoRoutine2()
	// presuffix.GoRoutine3()
	// presuffix.IntensePlay()
	// presuffix.BufferedCh()
	// presuffix.SignalClose()
	// presuffix.TestNumerousIdeas()
	// presuffix.Lazy()
	// network.SetupTCP1()
	// network.Call()
	// webserver.StartHelloWorldServer()
	// webserver.PollUrls()
	// webserver.WorkWithForm()
	// statistics.WorkWithForm()
	// wiki.ServeWiki()
	// reader.Reader()
	reader.SlothFacts()
}

// func failUsers() {

// 	// you can't catch errors that wilGorouines3l be generated immediately
// 	// at run time
// 	defer func() {
// 		if r := recover(); r != nil {
// 			err, _ := r.(error)
// 			fmt.Println("err: ", err)
// 		}
// 	}()
// 	// Create a value of type User from the users package using struct literal.
// 	// However, since password is unexported, it cannot be compiled:
// 	// - unknown users.User field 'password' in struct literalx
// 	u := users.User{
// 		Name: "Hoanh",
// 		ID:   101,

// 		// password: "xxxx",
// 	}

// 	fmt.Printf("User: %#v\n", u)
// }
