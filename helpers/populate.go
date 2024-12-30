package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"quote-generator-backend/config"
	"quote-generator-backend/models"

	"go.mongodb.org/mongo-driver/bson"
)


func PopulateQuotesFromFile(filePath string) {
	err := config.ConnectDB()
	if (err != nil){
		log.Println(err.Error())
	}
	db := config.DB

	// Get the database and collection
	collection := db.Collection("quotes")

	// Read quotes from file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Parse the JSON data
	var quotes []models.Quote
	if err := json.Unmarshal(data, &quotes); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Prepare documents for insertion
	var documents []interface{}
	for _, quote := range quotes {
		documents = append(documents, bson.M{
			"text":     quote.Text,
			"category": quote.Category,
		})
	}

	// Insert quotes into the collection
	_, err = collection.InsertMany(context.TODO(), documents)
	if err != nil {
		log.Fatalf("Failed to insert quotes: %v", err)
	}

	fmt.Println("Successfully added quotes to the database!")
}

func main(){
	filePath := "D:/past/year 4 semester 1/MCP/Project/Quote_Generator/helpers/quotes.json"
	PopulateQuotesFromFile(filePath)
}
