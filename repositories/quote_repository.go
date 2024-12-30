package repositories

import (
	"context"
	"quote-generator-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuoteRepository struct {
    Collection *mongo.Collection
}

func (r *QuoteRepository) AddQuote(quote models.Quote) error {
	_, err := r.Collection.InsertOne(context.TODO(), quote)
	return err
}

func (r *QuoteRepository) GetQuotesByCategory(category string) ([]models.Quote, error) {
    var quotes []models.Quote

    cursor, err := r.Collection.Find(context.TODO(), bson.M{"category": category})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var quote models.Quote
        if err := cursor.Decode(&quote); err != nil {
            return nil, err
        }
        quotes = append(quotes, quote)
    }

    return quotes, nil
}

func (r *QuoteRepository) GetRandomQuotes(limit int, categories ...string) ([]models.Quote, error) {
    var quotes []models.Quote

    // Build the match filter
    var matchFilter bson.M
    if len(categories) > 0 {
        // If categories are provided, filter by category
        matchFilter = bson.M{"category": bson.M{"$in": categories}}
    }

    // Create the aggregation pipeline
    pipeline := mongo.Pipeline{}

    if len(matchFilter) > 0 {
        // Add the match stage if a filter was specified
        pipeline = append(pipeline, bson.D{{"$match", matchFilter}})
    }

    // Add the sample stage to randomly select documents
    pipeline = append(pipeline, bson.D{{"$sample", bson.D{{"size", limit}}}})

    // Run the aggregation pipeline
    cursor, err := r.Collection.Aggregate(context.TODO(), pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    // Decode the results
    for cursor.Next(context.TODO()) {
        var quote models.Quote
        if err := cursor.Decode(&quote); err != nil {
            return nil, err
        }
        quotes = append(quotes, quote)
    }

    return quotes, nil
}
