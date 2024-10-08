package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	Owner              primitive.ObjectID   `bson:"owner" json:"owner"`
	CardName           string               `bson:"card_name" json:"card_name"`
	CardType           bool                 `bson:"card_type" json:"card_type"`
	CardNumber         string               `bson:"card_number" json:"card_number"`
	CardExpirationDate string               `bson:"card_expiration_date" json:"card_expiration_date"`
	CardCVV            string               `bson:"card_cvv" json:"card_cvv"`
	Transactions       []primitive.ObjectID `bson:"transactions" json:"transactions"`
}
