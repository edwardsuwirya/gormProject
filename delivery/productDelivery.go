package delivery

import (
	"github.com/edwardsuwirya/gormProject/config"
	"github.com/edwardsuwirya/gormProject/repository"
	"github.com/edwardsuwirya/gormProject/useCase"
	"gorm.io/gorm"
)

type ProductDelivery struct {
	productUseCase useCase.ProductUseCase
}

func NewProductDelivery(db *gorm.DB) *ProductDelivery {
	repo := repository.NewProductRepo(db)
	return &ProductDelivery{
		productUseCase: useCase.NewProductUseCase(repo),
	}
}

func (pd *ProductDelivery) Show() {
	//_ = productUseCase.RegisterProduct(&entity.Product{
	//	ProductName: "Gelas Plastik",
	//	ProductCode: "999",
	//	CategoryId:  "747adc32-b1e6-11ea-846d-5232048075cb",
	//})
	//_ = productUseCase.RegisterProduct(&entity.Product{
	//	ProductName: "Pot Bunga",
	//	ProductCode: "888",
	//	Category:entity.Category{
	//		CategoryName: "Perlengkapan Taman",
	//	},
	//})
	res := pd.productUseCase.GetProductCollection("1", "1", "product_name")
	for _, p := range res {
		config.AppConfig.Logger.WithField("ProductDelivery", "FindAllProductDelivery").Info(p.ToString())
	}
	//res := pd.productUseCase.GetProductById("54091fe0-5520-11ea-bb2b-9378803a9e60")
	//pretty.Println(res)

	countResult := pd.productUseCase.GetTotalProduct()
	config.AppConfig.Logger.WithField("ProductDelivery", "TotalProductDelivery").Info("Total ", *countResult)
}
