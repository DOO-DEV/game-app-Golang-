package migrate

import (
	"game-app/config"
	"game-app/repository/migrator"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	mgr := migrator.New(cfg.MySql)
	mgr.Up()
}

func New(cfg config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "migrate mysqlDB",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}
}
