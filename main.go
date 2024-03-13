package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wahyuazizi/EduQueasy/config"
)

func main() {
	db := config.NewDB()
	fmt.Println(db)
}
