package main

import "fmt"

func main() {
	totalMessages := 100
	comChannel := make(chan int)
	go sender(comChannel, totalMessages)
	reader(comChannel)
}

func reader(comChannel chan int) {
	for msg := range comChannel {
		fmt.Println("MESSAGE RECEIVED FROM CHANNEL: ", msg)
	}
}

func sender(channel chan int, totalMsgs int) {
	for i := 0; i < totalMsgs; i++ {
		channel <- i
	}
	close(channel)
}
