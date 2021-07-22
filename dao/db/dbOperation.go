package db

import (
	"satplan/common"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	log "github.com/sirupsen/logrus"
)

var satDb *gorm.DB
var err error

func Close() {
	sqlDB, _ := satDb.DB()
	sqlDB.Close()
}

func init() {
	var err error
	dbName := common.GetEnvValue("DB_NAME", "sat.db")

	// github.com/mattn/go-sqlite3
	satDb, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		panic(err)
	}

	log.Info("database inited successfully")
}
