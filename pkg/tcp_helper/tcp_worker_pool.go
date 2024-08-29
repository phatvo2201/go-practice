package tcp_helper

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println("scan ", p)
		wg.Done()
	}
}
