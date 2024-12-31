package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
    Username string               `json:"username" bson:"username"`
    Favorites []primitive.ObjectID `json:"favorites" bson:"favorites"`
    Limit int   `json:"limit" bson:"limit"`
    Categories []string `json:"categories" bson:"categories"`
}
