package main

import (
	"fmt"
	"github.com/edwardsuwirya/gormProject/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type dbInitialization struct {
	dbEngine       string
	dataSourceName string
}

func NewDbInitialization(c *config.Config) *dbInitialization {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.Db.DbUser, c.Db.DbPassword, c.Db.DbHost, c.Db.DbPort, c.Db.SchemaName)
	return &dbInitialization{c.Db.DbEngine, dataSourceName}
}

func (dbi *dbInitialization) InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(dbi.dbEngine, dbi.dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	db.LogMode(true)
	return db, nil
}
