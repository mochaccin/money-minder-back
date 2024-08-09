package repositories

import (
	"context"
	"money-minder/internal/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SpendRepo struct {
	MongoCollection *mongo.Collection
}

func (r *SpendRepo) InsertSpend(spend *types.Spend) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne((context.Background()), spend)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *SpendRepo) DeleteCard(spendID string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(spendID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}

	result, err := r.MongoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
