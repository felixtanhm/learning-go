package main

import (
	"fmt"
	"time"
)

func hello(num int, done chan bool) {
	fmt.Println("Hello world goroutine", num)
	time.Sleep(4 * time.Second)
	done <- true
}

func main() {
	chnl := make(chan bool, 4)
	for i := 1; i <= 4; i++ {
		go hello(i, chnl)
	}
	for i := 1; i <= 4; i++ {
		<-chnl
	}

	fmt.Println("main function")
	close(chnl)
}

// Since all goroutines are started concurrently, they will each wait 4 seconds, but the main function will wait for all of them to complete.
// This setup should take a little over 4 seconds to complete in total, rather than 16 seconds, as each goroutine is sleeping independently.
