package global

import (
	"path/filepath"
	"strings"
	"sys-service-scaffolding/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	// dns: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := "root:limumu1997@tcp(" + config.Config.DBIP + ":3306)/home_device?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func sqliteConnect() {
	var err error
	vr_etcdfs := filepath.Join(config.Config.DataPath, "vr_etcdfs.db")
	DB, err = gorm.Open(sqlite.Open(vr_etcdfs), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
