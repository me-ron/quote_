package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quote struct {
    ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Text     string             `json:"text" bson:"text"`
    Category string             `json:"category" bson:"category"`
}
