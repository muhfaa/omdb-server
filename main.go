package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muhfaa/omdb-server/config"
	grpcapi "github.com/muhfaa/omdb-server/controllers/grpc"
	httpapi "github.com/muhfaa/omdb-server/controllers/http"
	mysqlrepo "github.com/muhfaa/omdb-server/repository/mysql"
	omdbrepo "github.com/muhfaa/omdb-server/repository/omdb"
	"github.com/muhfaa/omdb-server/service"
)

// func init() {
// 	config.ReadConfig()
// }

func connectMysqlDB() *sql.DB {
	mysqlDSL := mysqlrepo.MysqlDBDSL{
		Username: config.GetConfig().MYSQLUsername,
		Password: config.GetConfig().MYSQLPassword,
		Host:     config.GetConfig().MYSQLHost,
		Port:     config.GetConfig().MYSQLPort,
		DBName:   config.GetConfig().MYSQLDBName,
	}

	fmt.Println("mysqlDSL:", config.GetConfig().MYSQLUsername)

	db, err := sql.Open("mysql", mysqlDSL.GetDSN())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(64)
	db.SetMaxIdleConns(64)
	return db
}

func main() {
	httpClient := http.DefaultClient

	mysqlRepo := mysqlrepo.NewMysqlDB(connectMysqlDB())
	omdbRepo := omdbrepo.NewOMDBRepo(httpClient, config.GetConfig().OMDBURL, config.GetConfig().OMDBKey)
	searchService := service.NewSearchService(omdbRepo, mysqlRepo)
	singleService := service.NewSingleService(omdbRepo)

	go httpapi.RunServer(config.GetConfig().HTTPPort, searchService, singleService)
	fmt.Println("http api running at :" + config.GetConfig().HTTPPort)

	go grpcapi.RunGRPCServer(config.GetConfig().GRPCPort, searchService, singleService)
	fmt.Println("grpc server running at :" + config.GetConfig().GRPCPort)

	select {}
}
