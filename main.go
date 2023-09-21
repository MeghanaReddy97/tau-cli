package main

import (
	"log"
	"os"

	"github.com/taubyte/tau-cli/cli"
	"github.com/taubyte/tau-cli/i18n"
)

func main() {
	err := cli.Run(os.Args...)
	if err != nil {
		log.Fatal(i18n.AppCrashed(err))
	}
}

// 1.This is where our CLI application begins. It follows the standard structure for Go code.
// 2.The main function starts and guides our command-line application.
// 3.It imports packages like fmt, log, and custom packages from the github.com/taubyte/tau repository.
//   a.fmt helps us with formatting and printing.
//   b.log keeps a record of important events.
//   c.We also have some custom packages from the github.com/taubyte/tau repository that add specific functionalities.
// 4.It generates a new instance of TauCliApp from the commands package and proceeds to execute it, supplying the command-line arguments as input.