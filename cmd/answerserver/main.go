package answerserver

import (
	"game-app/config"
	answergrpcserver "game-app/delivery/grpcserver/answer"
	"game-app/repository/mysql"
	mysqlanswer "game-app/repository/mysql/answer"
	"game-app/service/answerservice"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	MysqlRepo := mysql.New(cfg.MySql)
	answerRepo := mysqlanswer.New(MysqlRepo)
	answerSvc := answerservice.New(answerRepo)

	server := answergrpcserver.New(answerSvc)
	server.Start()
}

func New(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "answer-grpc-server",
		Short: "run presence grpc server",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}
}
