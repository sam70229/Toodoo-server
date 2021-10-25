package main

/*
import (
	"Toodoo/config"
	"Toodoo/logger"
	"Toodoo/server"
	"fmt"
	"os"

)

func main() {

	_ = config.Defaults(config.Config)

	config.InitLogger(config.Config)

	l, err := logger.New()

	l.Init(os.Stdout, os.Stdout, os.Stderr)

	s, err := server.New(config.Config)

	if err != nil {
		fmt.Println(err)
	}

	if err = s.Run(config.Config); err != nil {
		logger.Error.Fatal("Could not start server", "error", err)
	}
	select {}
}
*/

import (
    "os"
    "Toodoo/cmd/commands"
)


func main() {
    if err := commands.Run(os.Args[1:]); err != nil {
        os.Exit(1)
    }
}
