package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	name   string               `bson:"first_name" json:"first_name"`
	cards  []primitive.ObjectID `bson:"cards" json:"cards"`
	spends []primitive.ObjectID `bson:"spends" json:"spends"`
}
