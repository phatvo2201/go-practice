package tcp_helper

import (
	"sync"
)

func ScanWithGo() {
	var wg sync.WaitGroup
	ports := make(chan int, 100)

	for i := 0; i <= cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 0; i <= 1024; i++ {
		ports <- i
		wg.Add(1)

	}

	wg.Wait()
	close(ports)
}
