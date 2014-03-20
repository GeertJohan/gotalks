// +build OMIT

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// START OMIT
	findDir := os.Getenv("HOME") + "/Workspaces/Go/devpath/src/github.com/GeertJohan"
	cmdFind := exec.Command("find", ".", "-name", "*.go")
	cmdFind.Dir = findDir

	cmdWc := exec.Command("xargs", "wc", "-l")
	cmdWc.Dir = findDir
	cmdWc.Stdin, _ = cmdFind.StdoutPipe()
	cmdWc.Stdout = os.Stdout
	cmdWc.Start() // start wc (don't wait for it to finish)

	err := cmdFind.Run() // start find, and wait for it to finish
	if err != nil {
		fmt.Printf("Error finding *.go files: %s\n", err)
		return
	}

	err = cmdWc.Wait() // wait for wc to finish
	if err != nil {
		fmt.Printf("Error wc-ing *.go files: %s\n", err)
		return
	}
	// END OMIT
}
