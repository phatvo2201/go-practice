package tcp_helper

import (
	"io"
	"log"
	"net"
)

func Echo(conn net.Conn) {

	defer conn.Close()
	b := []byte{}
	for {
		size, err := conn.Read(b)
		if err == io.EOF {
			log.Println("Client disconnect")
			break
		}
		if err != nil {
			log.Println("Error when connect")
			break
		}
		log.Printf("Got %d from byte client", size)
		if _, err := conn.Write(b[:size]); err != nil {
			log.Fatalln("Unable to write data")
		}

	}
}
