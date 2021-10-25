package commands

import (
    "github.com/spf13/cobra"
    "Toodoo/config"
	"Toodoo/logger"
	"Toodoo/server"
	"fmt"
	"os"
)

var serverCmd = &cobra.Command{
            Use: "server",
            Short: "Start the api server",
            Long: `Start API`,
        }

func init() {
    RootCmd.AddCommand(serverCmd)
    RootCmd.RunE = serverCmdF
}

func serverCmdF(command *cobra.Command, args []string) error {
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

