package commands

import (
    "github.com/spf13/cobra"
	"go.uber.org/zap"

    "Toodoo/config"
	"Toodoo/logger"
	"Toodoo/server"
	
	"fmt"
)

var serverCmd = &cobra.Command{
            Use: "server",
            Short: "Start the api server",
            Long: `Start API`,
			RunE: serverCmdF,
        }

func init() {
    RootCmd.AddCommand(serverCmd)
}

func serverCmdF(command *cobra.Command, args []string) error {
    _ = config.Defaults(config.Config)

	logger.InitLogger(config.Config)

	s, err := server.New(config.Config)

	if err != nil {
		fmt.Println(err)
	}

	if err = s.Run(config.Config); err != nil {
		zap.S().Fatalw("error", "error", err)
	}
	select {}
}

