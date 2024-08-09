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

func (r *SpendRepo) DeleteSpend(spendID string) (*mongo.DeleteResult, error) {
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

func (r *SpendRepo) FindSpendByID(spendID string) (*types.Spend, error) {
	id, err := primitive.ObjectIDFromHex(spendID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", id}}

	var spend types.Spend

	err = r.MongoCollection.FindOne(context.Background(), filter).Decode(&spend)
	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &spend, nil
}

func (r *SpendRepo) GetSpendsByCardID(cardID string) ([]*types.Spend, error) {
	filter := bson.M{"card_id": cardID}

	var spends []*types.Spend

	cursor, err := r.MongoCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var spend types.Spend
		if err := cursor.Decode(&spend); err != nil {
			return nil, err
		}
		spends = append(spends, &spend)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return spends, nil
}

func (r *SpendRepo) GetSpendsByUserID(userID string) ([]*types.Spend, error) {
	filter := bson.M{"user_id": userID}

	var spends []*types.Spend

	cursor, err := r.MongoCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var spend types.Spend
		if err := cursor.Decode(&spend); err != nil {
			return nil, err
		}
		spends = append(spends, &spend)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return spends, nil
}
