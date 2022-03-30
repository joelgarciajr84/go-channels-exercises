package main

import (
	"fmt"
)

func main() {
	quitChannel := make(chan int)

	comChannel := gen(quitChannel)

	receive(comChannel, quitChannel)

	fmt.Println("Going exit ...")
}

func gen(quit <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(comChannel, quitChannel <-chan int) {
	for {
		select {

		case message, ok := <-comChannel:
			if ok {
				fmt.Println("Received from Communication Channel: ", message)
			}
		case message, ok := <-quitChannel:
			if ok {
				fmt.Println("Quit Channel just sent:", message)
			}
			return

		}
	}
}
