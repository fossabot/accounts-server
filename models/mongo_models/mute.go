package mongo_models

import (
	"github.com/UsagiBooru/accounts-server/gen"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoMuteStruct struct {
	// MongoのユニークID
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	// ミュートID
	MuteID int32 `json:"muteID,omitempty"`

	// ミュート種別
	TargetType string `json:"targetType,omitempty"`

	// 対象のタグ/絵師ID
	TargetID int32 `json:"targetID,omitempty"`
}

func (f *MongoMuteStruct) ToOpenApi() *gen.MuteStruct {
	resp := gen.MuteStruct{
		MuteID:     f.MuteID,
		TargetType: f.TargetType,
		TargetID:   f.TargetID,
	}
	return &resp
}
