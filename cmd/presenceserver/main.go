package presence_grpc_server

import (
	"game-app/adapter/redis"
	"game-app/config"
	"game-app/delivery/grpcserver/presenceserver"
	"game-app/repository/redis/redispresence"
	"game-app/service/presenceservice"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	redisAdapter := redis.New(cfg.Redis)

	presenceRepo := redispresence.New(redisAdapter)
	presenceSvc := presenceservice.New(presenceRepo, cfg.PresenceService)

	server := presenceserver.New(presenceSvc)
	server.Start()
}

func New(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "presence-grpc-server",
		Short: "run presence grpc server",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}
}
