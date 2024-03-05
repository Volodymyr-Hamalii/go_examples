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

func renderTestStr() {
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

func renderNestedItem() {
	data := map[string]interface{}{
		"itinerary": []map[string]interface{}{
			{
				"itinerary_items": []map[string]interface{}{
					{
						"location": map[string]interface{}{
							"city": "City1",
						},
						"name": "Item1",
					},
					{
						"location": map[string]interface{}{
							"city": "City2",
						},
						"name": "Item2",
					},
				},
			},
		},
	}

	// Adjust the template to render the name of the first item
	template := "{{#itinerary}}{{#itinerary_items}}{{#location}}{{#city}}{{name}}{{/city}}{{/location}}{{/itinerary_items}}{{/itinerary}}"


	result, err := mustache.Render(template, data)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return
	}

	fmt.Println(result)
}

func main() {
	renderNestedItem()
}
