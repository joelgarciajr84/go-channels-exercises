package main

import (
	"fmt"
)

func main() {
	engine := gen()
	receive(engine)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	newChannel := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			newChannel <- i
		}
		close(newChannel)
	}()

	return newChannel
}

func receive(channel <-chan int) {

	for v := range channel {
		fmt.Println("Value receive from channel is: ", v)
	}

}
