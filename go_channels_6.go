package main

import (
	"fmt"
)

func main() {

	communicationChannel := make(chan int)
	numberOfGoRoutines := 10

	go engine(numberOfGoRoutines, communicationChannel)
	readMessagesFromChannel(communicationChannel)

}

func engine(nGRoutines int, comChannel chan<- int) {
	numberOfMessages := 10
	for i := 0; i < nGRoutines; i++ {

		go func(numberOfMessages int) {
			for msg := 0; msg < numberOfMessages; msg++ {
				fmt.Println("ADDING MESSAGE", msg, "TO CHANNEL")
				comChannel <- msg
			}
			close(comChannel)

		}(numberOfMessages)

	}

}

func readMessagesFromChannel(channel chan int) {
	for v := range channel {
		fmt.Println("MESSAGE RECEIVED ", v)
	}
}
