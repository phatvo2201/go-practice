package main

import (
	"time"

	"github.com/phatvo2201/pkg/tcp_helper"
)

func main() {
	tcp_helper.ScanWithGo()
	time.Sleep(1 * time.Second)
}
