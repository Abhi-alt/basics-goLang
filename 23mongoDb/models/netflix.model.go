package netflixModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ObjectId primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie    string             `json:"movie,omitempty"`
	Watched  bool               `json:"watched,omitempty"`
}
