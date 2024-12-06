package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Spend struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Date        string             `bson:"date" json:"date"`
	Category    string             `bson:"category" json:"category"`
	Amount      int                `bson:"amount" json:"amount"`
	Owner       primitive.ObjectID `bson:"owner" json:"owner"`
	PaymentCard primitive.ObjectID `bson:"payment_card" json:"payment_card"`
}
