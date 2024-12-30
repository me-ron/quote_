package repositories

import (
	"context"
	"quote-generator-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

    var matchFilter bson.M
    if len(categories) > 0 {
        matchFilter = bson.M{"category": bson.M{"$in": categories}}
    }

    pipeline := mongo.Pipeline{}

    if len(matchFilter) > 0 {
        pipeline = append(pipeline, bson.D{{"$match", matchFilter}})
    }

    pipeline = append(pipeline, bson.D{{"$sample", bson.D{{"size", limit}}}})

    cursor, err := r.Collection.Aggregate(context.TODO(), pipeline)
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

func (r *QuoteRepository) GetAllCategories(id ...primitive.ObjectID) ([]string, error) {
    var categories []string

    var filter bson.M
    if len(id) > 0 && id[0] != primitive.NilObjectID {
        filter = bson.M{"user_id": bson.M{"$in": id[0]}}
    } else {
        filter = bson.M{}
    }

    cursor, err := r.Collection.Distinct(context.TODO(), "category", filter)
    if err != nil {
        return nil, err
    }

    for _, category := range cursor {
        categories = append(categories, category.(string))
    }

    return categories, nil
}
