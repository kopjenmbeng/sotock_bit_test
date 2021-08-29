package cmd

import (
	"github.com/kopjenmbeng/goconf"
	"github.com/kopjenmbeng/sotock_bit_test/internal/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(apiCommand)
}

var apiCommand = &cobra.Command{
	Use: "api",
	PreRun: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "apiCommand").Println("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "apiCommand").Println("Run done")
		api.NewServer(
			goconf.GetString("host.api_address"),
			logger,
			telemetry,
			Db.Read(),
			Db.Write(),
			api.JWE(jw),
		).Serve()
	},
}
