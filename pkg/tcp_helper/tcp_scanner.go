package tcp_helper

import (
	"log"
	"net"
)

func TcpScan() {

	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err != nil {
		log.Println("fail to connect")
	}
	log.Println("connect success !")
}
