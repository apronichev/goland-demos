package main

import (
	"encoding/json"
	jsonv2 "encoding/json/v2"
	"fmt"
	"log"
	"time"
)

// Product with various field types and tags
type Product struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Price     float64        `json:"price"`
	InStock   bool           `json:"in_stock"`
	Tags      []string       `json:"tags,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	Internal  string         `json:"-"` // Always omitted
	Metadata  map[string]any `json:"metadata,omitempty"`
}

func main() {
	// Build with: GOEXPERIMENT=jsonv2 go run main_test.go
	fmt.Println("=== JSON v2 Experimental Features ===\n")

	// Sample data
	product := Product{
		ID:        "prod-123",
		Name:      "Super Widget",
		Price:     29.99,
		InStock:   true,
		Tags:      []string{"gadget", "popular"},
		CreatedAt: time.Now(),
		Internal:  "secret-data",
		Metadata: map[string]any{
			"color":  "blue",
			"weight": 150,
		},
	}

	// 1. Using json/v2 package directly
	fmt.Println("1. JSON v2 marshaling:")
	v2Data, err := jsonv2.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n\n", v2Data)

	// 2. Standard json package (uses v2 implementation when GOEXPERIMENT=jsonv2)
	fmt.Println("2. Standard json package (powered by v2 internally):")
	standardData, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n\n", standardData)

	// 3. Unmarshal with v2
	fmt.Println("3. Unmarshaling with v2:")
	jsonData := `{
		"id": "prod-456",
		"name": "Another Widget",
		"price": 19.99,
		"in_stock": true,
		"tags": ["new", "sale"],
		"created_at": "2024-01-15T10:30:00Z"
	}`

	var newProduct Product
	err = jsonv2.Unmarshal([]byte(jsonData), &newProduct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Unmarshaled: %+v\n\n", newProduct)

	// 4. Error message comparison
	fmt.Println("4. Error handling comparison:")

	// Invalid JSON - price as string instead of number
	badJSON := `{"name": "Bad JSON", "price": "not-a-number"}`

	var p1, p2 Product
	err1 := json.Unmarshal([]byte(badJSON), &p1)
	err2 := jsonv2.Unmarshal([]byte(badJSON), &p2)

	fmt.Printf("Standard json error: %v\n", err1)
	fmt.Printf("json/v2 error: %v\n\n", err2)

	// 5. Performance demonstration
	fmt.Println("5. Performance benefits (decoding is faster):")

	// Create a large JSON array
	type Item struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	items := make([]Item, 10000)
	for i := range items {
		items[i] = Item{
			ID:    i,
			Name:  fmt.Sprintf("item-%d", i),
			Value: i * 100,
		}
	}

	// Marshal once
	largeJSON, _ := jsonv2.Marshal(items)
	fmt.Printf("JSON size: %d bytes\n", len(largeJSON))

	// Time standard json
	start := time.Now()
	var decoded1 []Item
	json.Unmarshal(largeJSON, &decoded1)
	standardTime := time.Since(start)

	// Time v2
	start = time.Now()
	var decoded2 []Item
	jsonv2.Unmarshal(largeJSON, &decoded2)
	v2Time := time.Since(start)

	fmt.Printf("Standard json decoded in: %v\n", standardTime)
	fmt.Printf("json/v2 decoded in: %v\n", v2Time)
	fmt.Printf("v2 is %.2fx faster\n\n", float64(standardTime)/float64(v2Time))

	// 6. Demonstrating consistent behavior
	fmt.Println("6. Consistent handling:")

	dupJSON := `{"name": "first", "name": "second", "name": "third"}`

	var dup1, dup2 map[string]string
	json.Unmarshal([]byte(dupJSON), &dup1)
	jsonv2.Unmarshal([]byte(dupJSON), &dup2)

	fmt.Printf("Standard json result: %v (silently uses last value)\n", dup1)
	fmt.Printf("json/v2 result: %v (consistent behavior)\n\n", dup2)
}
