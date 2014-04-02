// +build OMIT

package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
)

func main() {
	// START OMIT
	box, err := rice.FindBox("some-files")
	if err != nil {
		panic(err)
	}
	str := box.MustString("info.txt")
	fmt.Println(str)
	// END OMIT
}
