package cmd

import (
	"game-app/cmd/answerserver"
	"game-app/cmd/app"
	"game-app/cmd/game"
	"game-app/cmd/migrate"
	presence_grpc_server "game-app/cmd/presenceserver"
	"game-app/config"
	"game-app/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

// ExitFailure status code.
const ExitFailure = 1

func Execute() {
	cfg := config.New()

	root := &cobra.Command{
		Use:   "game-app",
		Short: "a quiz application that two players answers question in a game",
	}

	root.AddCommand(migrate.New(cfg))
	root.AddCommand(presence_grpc_server.New(cfg))
	root.AddCommand(app.New(cfg))
	root.AddCommand(game.New(cfg))
	root.AddCommand(answerserver.New(cfg))

	if err := root.Execute(); err != nil {
		logger.Logger.Error("failed to execute root command", zap.Error(err))
		os.Exit(ExitFailure)
	}
}
