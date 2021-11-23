package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"url_shortner/config"
	"url_shortner/internal/adapters/driven/mysql"
	"url_shortner/internal/adapters/driven/redis"
	"url_shortner/internal/adapters/driving/httphdl"
	"url_shortner/internal/core/service"
	"url_shortner/internal/repositories/urlRepo"
	"url_shortner/pkg/uidgen"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Shortner",
		Short: "shorting urls",
		Long:  "shorting long urls",
		Run: func(cmd *cobra.Command, args []string) {
			e := echo.New()
			mysqlConf,redisConf,serverConf := config.LoadConfigs()
			MYSQL := mysql.InitMysql(mysqlConf.Host,mysqlConf.Port,mysqlConf.UserName,mysqlConf.Password,mysqlConf.DbName)
			REDIS := redis.InitRedis(redisConf.Addr,redisConf.UserName,redisConf.Password,redisConf.DbName)
			dbrepo := urlRepo.NewDb(MYSQL)
			redrepo := urlRepo.NewCache(REDIS)
			urlservices := service.New(dbrepo,uidgen.New(),redrepo)
			hdl := httphdl.New(urlservices)
			e.POST("/new",hdl.Save)
			e.POST("/redirect",hdl.Read)

			e.Logger.Fatal(e.Start(serverConf.Host + ":" + serverConf.Port))
		},
	}
)

func Execute() error{
	return rootCmd.Execute()
}
