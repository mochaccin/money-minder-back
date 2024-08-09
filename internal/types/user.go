package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name     string               `bson:"name" json:"name"`
	Password string               `bson:"password" json:"password"`
	Cards    []primitive.ObjectID `bson:"cards" json:"cards"`
	Spends   []primitive.ObjectID `bson:"spends" json:"spends"`
}
