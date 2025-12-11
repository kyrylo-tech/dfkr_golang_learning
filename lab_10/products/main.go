package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Supplier struct {
	Name  string
	Phone string
}

type Product struct {
	ID       int
	Name     string
	Supplier Supplier
	Tags     []string
}

func loadProducts(path string) ([]Product, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','          // за замовчуванням
	r.FieldsPerRecord = -1 // дозволяємо різну кількість полів

	// читаємо заголовок
	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}

	var products []Product

	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if len(record) < 5 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		p := Product{
			ID:   id,
			Name: record[1],
			Supplier: Supplier{
				Name:  record[2],
				Phone: record[3],
			},
			Tags: strings.Split(record[4], "|"),
		}

		products = append(products, p)
	}

	return products, nil
}

func groupByTag(products []Product) map[string][]Product {
	result := make(map[string][]Product)

	for _, p := range products {
		for _, tag := range p.Tags {
			tag = strings.TrimSpace(tag)
			if tag == "" {
				continue
			}
			result[tag] = append(result[tag], p)
		}
	}

	return result
}

func main() {
	products, err := loadProducts("products.csv")
	if err != nil {
		fmt.Println("Error loading products:", err)
		return
	}

	groups := groupByTag(products)

	for tag, prods := range groups {
		fmt.Printf("=== Категорія: %s ===\n", tag)
		for _, p := range prods {
			fmt.Printf("- [%d] %s (постачальник: %s)\n", p.ID, p.Name, p.Supplier.Name)
		}
		fmt.Println()
	}
}
