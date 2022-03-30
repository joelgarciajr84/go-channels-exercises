package main

import "fmt"

func main() {
	comChannel := make(chan int)
	go sender(comChannel)
	reader(comChannel)
}

func reader(comChannel chan int) {
	for msg := range comChannel {
		fmt.Println("MESSAGE RECEIVED FROM CHANNEL: ", msg)
	}
}

func sender(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}
