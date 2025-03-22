package main

import (
	grpc "berezovskyivalerii/client-rpc-csv/internal/grpc"
	"log"
	"fmt"
	//"github.com/berezovskyivalerii/server-rpc-csv/pkg/domain"
)

func main() {
	client, err := grpc.NewClient(9000)
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}
	defer client.CloseConnection()

	//url := "http://localhost:8000/products?format=csv"

	// Параметры запроса List
	// page := int32(1)
	// pageSize := int32(6)
	// sortField := "name"
	// sortOrder := "asc"

	// // Вызов List
	// resp, err := client.List(page, pageSize, sortField, sortOrder)
	// if err != nil {
	// 	log.Fatal("List failed:", err)
	// }

	// fmt.Println("Total products:", resp.TotalProducts)
	// for _, p := range resp.Products {
	// 	fmt.Printf("Name: %s | Price: %.2f | Changes: %d | Updated: %s\n",
	// 		p.ProductName, p.Price, p.PriceChangeCount, p.LastUpdated)
	// }

	//Вызов запроса Fetch
	url := "http://localhost:8000/products?format=csv"
	resp, err := client.Fetch(url)
	if err != nil {
		log.Fatal("Fetch failed:", err)
	}

	fmt.Println("Fetch Response:", resp.Success, resp.Message)
}