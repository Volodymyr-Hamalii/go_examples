package main

import (
	"encoding/json"
	"fmt"
	"github.com/cbroglie/mustache"
)

// Product struct definition
type Product struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	// Creating an instance of the Product struct
	product := Product{
		Name:    "Some product",
		Address: "Some street",
	}

	// Convert struct to map
	productMap := make(map[string]interface{})
	productJSON, _ := json.Marshal(product)
	json.Unmarshal(productJSON, &productMap)

	fmt.Println("productMap\n", productMap)

	// Template string
	templateString := "Product Name: {{name}}, Address: {{address}}"

	// Render the template
	renderedString, err := mustache.Render(templateString, productMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the rendered template
	fmt.Println(renderedString)
}
