package mongo_models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoSequence struct {
	// MongoのユニークID
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	// シーケンスのキー
	Key string `json:"key" bson:"key" validate:"alphanum,min=1,max=32"`

	// シーケンスの現在値
	Value int32 `json:"value" bson:"value" validate:"gte=0"`
}
