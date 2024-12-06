package types

type User struct {
	Name     string  `bson:"name" json:"name"`
	Password string  `bson:"password" json:"password"`
	Balance  int     `bson:"balance,omitempty" json:"balance,omitempty"`
	Cards    []Card  `bson:"cards,omitempty" json:"cards,omitempty"`
	Spends   []Spend `bson:"spends,omitempty" json:"spends,omitempty"`
}
