// +build OMIT

package main

import (
	"fmt"
	"github.com/GeertJohan/go.linenoise"
)

func main() {
	// START OMIT
	line, err := linenoise.Line("Say something: ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("You said: %s\n", line)
	// END OMIT
}
