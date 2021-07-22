package db

import (
	"satplan/common"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var db *gorm.DB
var err error
var configParams = make(map[string]map[string]string)

func Close() {
	//db.Close()
}

func init() {
	var err error
	dbName := common.GetEnvValue("DB_NAME", "sat.db")

	// github.com/mattn/go-sqlite3
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Info("database inited successfully")
}
