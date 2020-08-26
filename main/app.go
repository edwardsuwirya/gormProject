package main

import (
	"github.com/edwardsuwirya/gormProject/config"
	"github.com/edwardsuwirya/gormProject/repository"
	"github.com/edwardsuwirya/gormProject/useCase"
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
	db, err := NewDbInitialization(a.appConfig).InitDB()
	if err != nil {
		panic(err)
	}
	repo := repository.NewProductRepo(db)
	appUseCase := useCase.NewProductUseCase(repo)
	//_ = appUseCase.RegisterProduct(&entity.Product{
	//	ProductName: "Gelas Plastik",
	//	ProductCode: "999",
	//	CategoryId:  "747adc32-b1e6-11ea-846d-5232048075cb",
	//})
	//_ = appUseCase.RegisterProduct(&entity.Product{
	//	ProductName: "Pot Bunga",
	//	ProductCode: "888",
	//	Category:entity.Category{
	//		CategoryName: "Perlengkapan Taman",
	//	},
	//})
	//res:=appUseCase.GetProductCollection()
	//for _, p := range res {
	//	config.AppConfig.Logger.Info(p.ToString())
	//}
	res:= appUseCase.GetProductById("54091fe0-5520-11ea-bb2b-9378803a9e60")
	config.AppConfig.Logger.Info(res.ToString())

	//countResult := appUseCase.GetTotalProduct()
	//config.AppConfig.Logger.Info("Total ",*countResult)

}
func newApp() *app {
	conf := config.NewConfig()
	return &app{appConfig: conf}
}

func main() {
	newApp().run()
}
