// +build OMIT

package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

//START OMIT
var options struct {
	Verbose  bool   `short:"v" long:"version" description:"Enable verbose logging"`
	HTTPPort string `long:"http-port" default:"8080" description:"HTTP port to use"`
}

func main() {
	args, err := flags.Parse(&options) // uses os.Args (args given to command when executed)
	if err != nil {
		return
	}
	fmt.Println(options.Verbose)
	fmt.Println(options.HTTPPort)
	fmt.Println(args) // print any arguments given after the last option/flag
}

//END OMIT
