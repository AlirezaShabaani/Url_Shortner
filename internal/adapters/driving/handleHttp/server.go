package handleHttp

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"url_shortner/config"
	"url_shortner/internal/adapters/driven/mysql"
	"url_shortner/internal/adapters/driven/redis"
	"url_shortner/internal/core/service"
	"url_shortner/internal/repositories/urlRepo"
	"url_shortner/pkg/uidgen"
)

func StartServer()  {


	e := echo.New()
	mysqlConf, redisConf, serverConf := config.LoadConfigs()
	MYSQL := mysql.InitMysql(mysqlConf.Host, mysqlConf.Port, mysqlConf.UserName, mysqlConf.Password, mysqlConf.DbName)
	REDIS := redis.InitRedis(redisConf.Addr, redisConf.UserName, redisConf.Password, redisConf.DbName)
	dbrepo := urlRepo.NewDb(MYSQL)
	redrepo := urlRepo.NewCache(REDIS)
	urlservices := service.New(dbrepo, uidgen.New(), redrepo)
	hdl := New(urlservices) //hdl = handler
	e.POST("/new", hdl.Save)
	e.POST("/redirect", hdl.Read)


	// Graceful shutdown
	go func() {
		if err := e.Start(serverConf.Host +":" + serverConf.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
