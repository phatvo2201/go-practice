package main

import (
	"log"
	"net"
	"os"

	"github.com/phatvo2201/pkg/tcp_helper"
)

func main() {
	// Listen on local port 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	b := []byte{}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		os.Stdout.Write(b)
		go tcp_helper.ProxyHandler(conn)
	}
}

// func main() {
// 	listener, err := net.Listen("tcp", ":30010")
// 	if err != nil {
// 		log.Fatalln("Unable to bind to port")
// 	}
// 	log.Println("Listening on 0.0.0.0:30010")
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Fatalln("fail when connect to port")

// 		}
// 		go tcp_helper.Echo(conn)
// 	}

// }
