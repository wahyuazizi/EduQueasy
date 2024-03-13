package config

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", `root:password@tcp(localhost:3306)/EduQueasy_db`)
	if err != nil {
		panic(err)
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
