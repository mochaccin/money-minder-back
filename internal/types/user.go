package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name     string               `bson:"name" json:"name"`
	Password string               `bson:"password" json:"password"`
	Cards    []primitive.ObjectID `bson:"cards,omitempty" json:"cards,omitempty"`
	Spends   []primitive.ObjectID `bson:"spends,omitempty" json:"spends,omitempty"`
}
