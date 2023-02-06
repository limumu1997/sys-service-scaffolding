package global

import (
	"database/sql"
	"path/filepath"
	"strings"
	"sys-service-scaffolding/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MYSQL_DB *sql.DB
var SQLITE_DB *gorm.DB

func init() {
	if !strings.EqualFold(config.Config.DataPath, "") {
		sqliteConnect()
	}
	if !strings.EqualFold(config.Config.DBIP, "") {
		mysqlConnect()
	}
}

func mysqlConnect() {
	var err error
	dns := "exitpassMYSQL_DB2user:kRaq4FQsnNWyM9NFxvYUjl5FmEGY7uMC@tcp(" + config.Config.DBIP + ":3306)/exitpassMYSQL_DB2"
	MYSQL_DB, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	MYSQL_DB.SetConnMaxLifetime(60 * time.Second)
	MYSQL_DB.SetMaxOpenConns(100)
	MYSQL_DB.SetMaxIdleConns(16)
}

func sqliteConnect() {
	var err error
	vr_etcdfs := filepath.Join(config.Config.DataPath, "vr_etcdfs.db")
	SQLITE_DB, err = gorm.Open(sqlite.Open(vr_etcdfs), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
