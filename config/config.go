package config

import (
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type mysqlConf struct {
	Port     string
	Host     string
	UserName string
	Password string
	DbName   string
}

type serveConf struct {
	Port string
	Host string
}

type redisConf struct {
	Addr     string
	UserName string
	Password string
	DbName   int
}

func LoadConfigs() (mysqlConf *mysqlConf,redisConf *redisConf,serverConf *serveConf) {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	mysqlConf = loadSDBConfig(v)
	redisConf = loadRedisConf(v)
	serverConf = loadServeConfig(v)
	return
}

func loadSDBConfig(v *viper.Viper) *mysqlConf {
	port := v.GetString("MYSQL_PORT")
	host := v.GetString("MYSQL_HOST")
	user := v.GetString("MYSQL_USER")
	pass := v.GetString("MYSQL_PASS")
	DbName := v.GetString("MYSQL_DbName")
	if port == "" || host == "" || user == "" || pass == "" || DbName == "" {
		log.Fatal("db initialize failed because of empty params in connection address, check config")
	}
	return &mysqlConf{
		Port:     port,
		Host:     host,
		UserName: user,
		Password: pass,
		DbName:   DbName,
	}
}

func loadRedisConf(v *viper.Viper) *redisConf {
	port := v.GetString("REDIS_PORT")
	host := v.GetString("REDIS_HOST")
	user := v.GetString("REDIS_USER")
	pass := v.GetString("REDIS_PASS")
	DbName := v.GetString("REDIS_DbName")
	if port == "" || host == "" {
		log.Fatal("db initialize failed because of empty params in connection address, check config")
	}
	dbInt,err := strconv.Atoi(DbName)
	if err != nil {
		log.Println("can't parse string value of redis db name to integer")
		log.Fatal(err.Error())
	}
	return &redisConf{
		Addr:     host + ":" + port,
		UserName: user,
		Password: pass,
		DbName:   dbInt,
	}
}

func loadServeConfig(v *viper.Viper) *serveConf {
	port := v.GetString("SERVER_PORT")
	host := v.GetString("SERVER_HOST")
	if port == "" || host == "" {
		log.Fatal("service initialize failed because of empty params in connection address, check config")
	}
	return &serveConf{
		Port: port,
		Host: host,
	}
}

