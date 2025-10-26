package market

import (
	"fmt"
	"strings"

	"github.com/mazen160/go-random"
)

type Market struct {
	Products map[string]Product
}

type Product struct {
	Name   string
	Price  float64
	Weight float64
	Code   string
}

func NewMarket() *Market {
	return &Market{
		Products: make(map[string]Product), // Initialize the map here
	}
}

func (m *Market) AddProduct(Name string, Price float64, Weight float64) {
	CodeData, err := random.String(5)

	if err != nil {
		fmt.Println(err)
	}

	m.Products[strings.ToUpper(CodeData)] = Product{
		Name:   Name,
		Price:  Price,
		Weight: Weight,
		Code:   strings.ToUpper(CodeData),
	}
}

func (m *Market) GetProductInfo(p Product) {
	fmt.Printf("\nInfo about product:\n - Name: %s\n - Price: %.f\n - Weight: %.f\n - Code: %s\n\n",
		p.Name, p.Price, p.Weight, p.Code)
}

func (m *Market) GetProductByCode(code string) Product {
	return m.Products[code]
}

func (m *Market) GetProductsDialog() {
	fmt.Printf("\nAll products:")

	var code string

	for code := range m.Products {
		fmt.Printf("\n- %s", code)
	}

	fmt.Printf("\n\nType prefer code of product you need:")
	fmt.Scanln(&code)

	product := m.GetProductByCode(code)
	m.GetProductInfo(product)
}
