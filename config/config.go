package config

import (
	"sync"

	"github.com/spf13/viper"
)

type AppConfig struct {
	HTTPPort      string `mapstructure:"HTTP_Port"`
	GRPCPort      string `mapstructure:"GRPC_Port"`
	OMDBKey       string `mapstructure:"OMDB_Key"`
	OMDBURL       string `mapstructure:"OMDB_URL"`
	MYSQLUsername string `mapstructure:"MYSQL_Username"`
	MYSQLPassword string `mapstructure:"MYSQL_Password"`
	MYSQLHost     string `mapstructure:"MYSQL_Host"`
	MYSQLPort     string `mapstructure:"MYSQL_Port"`
	MYSQLDBName   string `mapstructure:"MYSQL_DBName"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	//re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	defaultConfig.HTTPPort = "8080"
	defaultConfig.GRPCPort = "8081"
	defaultConfig.OMDBKey = "faf7e5bb&s"
	defaultConfig.OMDBURL = "http://www.omdbapi.com/"
	defaultConfig.MYSQLUsername = "root"
	defaultConfig.MYSQLPassword = "password"
	defaultConfig.MYSQLHost = "127.0.0.1"
	defaultConfig.MYSQLPort = "3306"
	defaultConfig.MYSQLDBName = "omdb"

	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.BindEnv("HTTP_Port")
	viper.BindEnv("GRPC_Port")
	viper.BindEnv("OMDB_Key")
	viper.BindEnv("OMDB_URL")
	viper.BindEnv("MYSQL_Username")
	viper.BindEnv("MYSQL_Password")
	viper.BindEnv("MYSQL_Host")
	viper.BindEnv("MYSQL_Port")
	viper.BindEnv("MYSQL_DBName")

	err := viper.ReadInConfig()
	if err != nil {
		return &defaultConfig
	}

	var finalConfig AppConfig
	err = viper.Unmarshal(&finalConfig)
	if err != nil {
		return &defaultConfig
	}

	return &finalConfig
}
