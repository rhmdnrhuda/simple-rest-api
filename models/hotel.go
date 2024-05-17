package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Stars int                `json:"stars" bson:"stars"`
}
