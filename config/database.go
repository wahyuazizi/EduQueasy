package config

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type DatabaseConfig struct {
	DriverName string
	Host       string
	Port       int
	Username   string
	Password   string
	DbName     string
}

func NewDB() *sql.DB {

	// Inisialisasi VIPER
	config := viper.New()
	config.SetConfigFile("config.yaml")

	// Binding ke struct
	var dbConfig DatabaseConfig
	err := viper.UnmarshalKey("database", &dbConfig)
	if err != nil {
		panic(err.Error())
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
	db, err := sql.Open(dbConfig.DriverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db

	//migrate -database "mysql://root:password@tcp(localhost:3306)/EduQueasy" -path db/migrations up
	//migrate -database "mysql://root:password@tcp(localhost:3306)/EduQueasy" -path db/migrations version
	//migrate -database "mysql://root:password@tcp(localhost:3306)/EduQueasy" -path db/migrations force 20240309081214
	// migrate create -ext sql -dir db/migrations create_table_third
}
