package global

import (
	"database/sql"
	"sys-service-scaffolding/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	dns := "exitpassdb2user:kRaq4FQsnNWyM9NFxvYUjl5FmEGY7uMC@tcp(" + config.Config.DbIP + ":3306)/exitpassdb2"
	DB, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(60 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
}
