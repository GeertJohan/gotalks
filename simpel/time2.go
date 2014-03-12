// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	verjaardag, _ := time.Parse("Jan 2 2006", "Mar 2 2014") // time.Time
	leeftijd := time.Since(verjaardag)                      // time.Duration
	fmt.Printf("Go 1.2.1 is %d dagen oud.\n", leeftijd/(time.Hour*24))
	// END OMIT
}
