package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Spend struct {
	name        string             `bson:"name" json:"name"`
	date        string             `bson:"date" json:"date"`
	category    string             `bson:"category" json:"category"`
	amount      int                `bson:"amount" json:"amount"`
	owner       primitive.ObjectID `bson:"owner" json:"owner"`
	paymentCard primitive.ObjectID `bson:"payment_card" json:"payment_card"`
}
