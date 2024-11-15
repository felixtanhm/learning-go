// Activity 1
// "Received 0 true"
// "Received 1 true"
// "Received 2 true"
// "Received 3 true"
// "Received 4 true"
// "Received 5 true"
// "Received 6 true"
// "Received 7 true"
// "Received 8 true"
// "Received 9 true"

// Activity 2
// "1"
// "2"
// The program exits itself with an exit status of 2 due to the channel having a specified buffer of only 2.

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// fmt.Println(<-ch)

	ch <- 1
	// ch <- 2
	// fmt.Println(<-ch)

	// // ch <- 4
	// ch <- 5
	// fmt.Println(<-ch)
	fmt.Println(<-ch)

}
