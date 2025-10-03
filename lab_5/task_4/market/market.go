package market

import (
    "github.com/mazen160/go-random"
    "fmt"
    "strings"
    "errors"
)

type Market struct {
    Products []Product
}

type Product struct {
	Name string
	Price float64
	Weight float64
    Code string
}

func (m *Market) AddProduct(Name string, Price float64, Weight float64) {
    CodeData, err := random.String(5)
    
    if err != nil {
        fmt.Println(err)
    }
    
    m.Products = append(m.Products, Product{
        Name: Name,
        Price: Price,
        Weight: Weight,
        Code: strings.ToUpper(CodeData),
    })
}

func (m *Market) GetProductInfo(p Product) {
    fmt.Printf("\nInfo about product:\n - Name: %s\n - Price: %.f\n - Weight: %.f\n - Code: %s\n\n",
        p.Name, p.Price, p.Weight, p.Code)
}

func (m *Market) GetProductByCode(code string) (Product, error) {
    for i := 0; i < len(m.Products); i++ {
        if m.Products[i].Code == code {
            return m.Products[i], nil
        }
    }

    return Product{}, errors.New("Product with this code not found")
}

func (m *Market) GetProductsDialog() {
    fmt.Printf("\nAll products:")
    
    var code string

    for i := 0; (i < len(m.Products)); i++ {
        fmt.Printf("\n- %s", m.Products[i].Code)
    }
    
    fmt.Printf("\n\nType prefer code of product you need:")
    fmt.Scanln(&code)

    product, err := m.GetProductByCode(code)
    if (err != nil) {
        fmt.Printf("\n%s\n\n", err)
        return
    }

    m.GetProductInfo(product)
}