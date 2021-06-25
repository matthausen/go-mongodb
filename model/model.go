package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type (
	Note struct {
		ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Text string `json:"text"`
		Archived bool `json:"archived"`
	}
)
