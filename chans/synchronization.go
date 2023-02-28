package chans

import (
	"fmt"
	"time"
)

// goroutine, just waiting
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

// 3rd menu "W = Wait for goroutine"
func ChanSync() {
	// create a boolean channel
	done := make(chan bool, 1)
	go worker(done)
	// waiting for worker
	<-done
}
