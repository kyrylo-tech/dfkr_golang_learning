package market

// Define the Market struct
type Market struct {
    Products []Product
}

// Method to add a product to the market
func (m *Market) AddProduct(p Product) {
    m.Products = append(m.Products, p)
}
