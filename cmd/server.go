package cmd

import (
	"github.com/spf13/cobra"
	"url_shortner/config"
	"url_shortner/internal/adapters/driven/mysql"
	"url_shortner/internal/adapters/driven/redis"
	"url_shortner/internal/adapters/driving/httphdl"
	"url_shortner/internal/core/service"
	"url_shortner/internal/repositories/urlservices"
	"url_shortner/pkg/uidgen"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Shortner",
		Short: "shorting urls",
		Long:  "shorting long urls",
		Run: func(cmd *cobra.Command, args []string) {
			mysqlConf,redisConf,serverConf := config.LoadConfigs()
			MYSQL := mysql.InitMysql(mysqlConf.Host,mysqlConf.Port,mysqlConf.UserName,mysqlConf.Password,mysqlConf.DbName)
			REDIS := redis.InitRedis(redisConf.Addr,redisConf.UserName,redisConf.Password,redisConf.DbName)
			dbrepo := urlservices.NewDb(MYSQL)
			redrepo := urlservices.NewCache(REDIS)
			urlservices := service.New(dbrepo,uidgen.UIDGen().New(),redrepo)
			hdl := httphdl.New(urlservices)

		},
	}
)

func Execute() error{
	return rootCmd.Execute()
}
