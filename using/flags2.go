// +build OMIT

package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

// START OMIT
var options struct {
	Verbose bool `long:"verbose" short:"v" description:"Show verbose debug information"`

	Add struct {
		Force       bool `short:"f" long:"force" description:"Allow adding otherwise ignored files."`
		Interactive bool `short:"i" long:"interactive" description:"Add modified contents in the working tree interactively to the index. Optional path arguments may be supplied to limit operation to a subset of the working tree. See “Interactive mode” for details."`
	} `command:"add"`

	Commit struct {
		All     bool   `short:"a" long:"all" description:"Tell the command to automatically stage files that have been modified and deleted, but new files you have not told git about are not affected."`
		Message string `short:"m" long:"message" description:"Use the given <msg> as the commit message."`
	} `command:"commit"`

	Push struct {
	} `command:"push"`
}

// END OMIT

func main() {
	args, err := flags.Parse(&options)
	if err != nil {
		return
	}
	fmt.Println(args)
}
