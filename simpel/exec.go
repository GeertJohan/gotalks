// +build OMIT

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// START OMIT
	cmd := exec.Command("git", "log", "-5", "--pretty=format:%h%x09%an%x09%s")
	cmd.Dir = "/home/geertjohan/.oh-my-zsh"
	out, err := cmd.CombinedOutput()
	fmt.Println(err)
	fmt.Println(string(out))
	// END OMIT
}
