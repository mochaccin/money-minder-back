package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
	id                 primitive.ObjectID
	owner              primitive.ObjectID   `bson:"owner" json:"owner"`
	cardName           string               `bson:"card_name" json:"card_name"`
	cardType           bool                 `bson:"card_type" json:"card_type"`
	cardNumber         string               `bson:"card_number" json:"card_number"`
	cardExpirationDate string               `bson:"card_expiration_date" json:"card_expiration_date"`
	cardCVV            string               `bson:"card_cvv" json:"card_cvv"`
	transactions       []primitive.ObjectID `bson:"transactions" json:"transactions"`
}
