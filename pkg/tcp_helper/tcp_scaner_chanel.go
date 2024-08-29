package tcp_helper

import (
	"fmt"
	"net"
	"sync"
)

func ScanWithGo() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			defer wg.Done()
			wg.Add(1)
			add := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", add)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open \n", j)

		}(i)
	}
	wg.Wait()
}
