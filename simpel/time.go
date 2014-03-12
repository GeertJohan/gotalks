// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	if time.Now().Hour() < 12 {
		fmt.Println("Goeie morgen.")
	} else {
		fmt.Println("Goeie middag (of avond).")
	}
	// END OMIT
}
