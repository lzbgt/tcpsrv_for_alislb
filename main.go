package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/droundy/goopt"
)

var server = goopt.String([]string{"-s", "--server"}, "0.0.0.0", "Server to connect to (and listen if listening)")
var port = goopt.Int([]string{"-p", "--port"}, 19822, "Port to connect to (and listen to if listening)")

var listen = goopt.Flag([]string{"-l", "--listen"}, []string{}, "Create a listening TFO socket", "")

func main() {
	goopt.Parse(nil)

	// Listen for incoming connections.
	l, err := net.Listen("tcp", *server+":"+strconv.Itoa(*port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + *server + ":" + strconv.Itoa(*port))
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 20)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		//fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("\n"))
	// Close the connection when you're done with it.
	conn.Close()
}
