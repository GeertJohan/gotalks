// +build OMIT

package main

import (
	"log"
	"strconv"
)

// START OMIT
func producer(messageCh chan string) {
	for i := 0; i < 5; i++ {
		message := "bericht " + strconv.Itoa(i)
		messageCh <- message // HL
	}
}

func printer(messageCh chan string) {
	for {
		message := <-messageCh // HL
		log.Println(message)
	}
}

func main() {
	messageCh := make(chan string)

	go printer(messageCh)
	producer(messageCh)
}

// END OMIT
