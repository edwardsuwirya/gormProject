package main

import (
	"flag"
	"fmt"
	"github.com/edwardsuwirya/gormProject/config"
	"github.com/edwardsuwirya/gormProject/delivery"
	"gorm.io/gorm"
	"strings"
)

const (
	AppName    = "Enigma GORM Example Project"
	AppTagLine = "Example GO Project using GORM - MySQL"
	Version    = "0.0.1"
)

type app struct {
	appConfig *config.Config
	db        *gorm.DB
}

func (a app) run() {
	//cat := delivery.NewCategoryDelivery(db)
	//cat.Show()
	prod := delivery.NewProductDelivery(a.db)
	prod.Show()

}

func newApp() *app {
	fmt.Println(AppName)
	fmt.Println(AppTagLine)
	fmt.Printf("%s\n", strings.Repeat("=", 50))
	conf := config.NewConfig()
	db, err := NewDbInitialization(conf).InitDB()
	if err != nil {
		panic(err)
	}
	return &app{appConfig: conf, db: db}
}

func main() {
	migrationPtr := flag.Bool("migrate", false, "Run database migration")
	flag.Parse()
	if *migrationPtr {
		newApp().runMigration()
	} else {
		newApp().run()
	}

}
