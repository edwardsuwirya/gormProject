package main

import (
	"fmt"
	"github.com/edwardsuwirya/gormProject/config"
	"github.com/edwardsuwirya/gormProject/delivery"
	"strings"
)

const (
	AppName    = "Enigma GORM Example Project"
	AppTagLine = "Example GO Project using GORM - MySQL"
	Version    = "0.0.1"
)

type app struct {
	appConfig *config.Config
}

func (a app) run() {
	fmt.Println(AppName)
	fmt.Println(AppTagLine)
	fmt.Printf("%s\n", strings.Repeat("=", 50))
	db, err := NewDbInitialization(a.appConfig).InitDB()
	if err != nil {
		panic(err)
	}
	//cat := delivery.NewCategoryDelivery(db)
	//cat.Show()
	prod := delivery.NewProductDelivery(db)
	prod.Show()

}
func newApp() *app {
	conf := config.NewConfig()
	return &app{appConfig: conf}
}

func main() {
	newApp().run()
}
