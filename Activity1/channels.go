package main

import "fmt"

func producer(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}

	close(channel)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
		// output be like:
		// Received 0 true
		// Received 1 true
		// ....
		// Received 9 true

	}
}
