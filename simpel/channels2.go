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
		messageCh <- message
	}
	close(messageCh) // HL
}

func printer(messageCh chan string) {
	for {
		message, ok := <-messageCh
		if !ok { // HL
			return // HL
		} // HL
		log.Println(message)
	}
}

func main() {
	messageCh := make(chan string)

	go producer(messageCh)
	printer(messageCh)
}

// END OMIT
