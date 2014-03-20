// +build OMIT

package main

import (
	"log"
	"time"
)

// START1 OMIT
import "math/rand"

// END1 OMIT

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START2 OMIT
func ping(pingChannel chan bool) {
	time.Sleep(time.Duration(rand.Int31n(200)) * time.Millisecond)
	pingChannel <- true
}

func main() {
	alphaCh := make(chan bool)
	go ping(alphaCh)
	betaCh := make(chan bool)
	go ping(betaCh)

	select {
	case <-alphaCh:
		log.Println("Alpha was first")
	case <-betaCh:
		log.Println("Beta was first")
	}
}

// END2 OMIT
