package tcp_helper

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var desAddress = "qlack.com.vn"

type Reader struct{}

func (rd *Reader) Read(b []byte) (int, error) {
	fmt.Println("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func ProxyHandler(src net.Conn) {

	des, err := net.Dial("tcp", desAddress)
	if err != nil {
		log.Println("fail when connect to destination service ...")
		return
	}
	go func() {
		size, err := io.Copy(des, src)
		if err != nil {
			log.Fatal()
		}
		log.Printf("forward %d to destination", size)

	}()
	if _, err := io.Copy(src, des); err != nil {
		log.Println("fail when send output back to source service ...")
		return

	}

}
