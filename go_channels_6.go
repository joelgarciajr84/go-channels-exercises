package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	communicationChannel := make(chan int)
	numberOfGoRoutines := 10
	wg.Add(numberOfGoRoutines)
	go engine(numberOfGoRoutines, communicationChannel)
	go readMessagesFromChannel(communicationChannel)
	wg.Wait()

}

func engine(nGRoutines int, comChannel chan<- int) {
	numberOfMessages := 10
	for i := 0; i < nGRoutines; i++ {

		go func(numberOfMessages int) {
			for msg := 0; msg < numberOfMessages; msg++ {
				fmt.Println("ADDING MESSAGE", msg, "TO CHANNEL")
				comChannel <- msg
			}
			wg.Done()

		}(numberOfMessages)

	}

}

func readMessagesFromChannel(channel chan int) {
	for v := range channel {
		fmt.Println("MESSAGE RECEIVED ", v)
	}
	// close(channel)
	wg.Done()

}
