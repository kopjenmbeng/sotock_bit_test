package cmd

import (
	"github.com/kopjenmbeng/goconf"
	"github.com/kopjenmbeng/sotock_bit_test/internal/grpc"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(grpcCommand)
}

var grpcCommand = &cobra.Command{
	Use: "grpc",
	PreRun: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "apiCommand").Println("PreRun done")
	},
	Run: func(cmd *cobra.Command, args []string) {
		defer logger.WithField("component", "apiCommand").Println("Run done")
		grpc.NewServer(logger,goconf.GetString("host.grpc_address")).Serve()
		// grpc.NewServer(
		// 	goconf.GetString("host.grpc_address"),
		// 	logger,
		// 	telemetry,
		// 	Db.Read(),
		// 	Db.Write(),
		// 	api.JWE(jw),
		// ).Serve()
	},
}
