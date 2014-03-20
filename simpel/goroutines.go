// +build OMIT

package main

import (
	"log"
	"time"
)

func printAfter(d time.Duration, message string) {
	time.Sleep(d)
	log.Println(message)
}

func main() {
	go printAfter(2*time.Second, "a")
	go printAfter(1*time.Second, "b")

	printAfter(3*time.Second, "klaar")
}
