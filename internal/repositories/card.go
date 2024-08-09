package repositories

import (
	"context"
	"fmt"
	"money-minder/internal/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardRepo struct {
	MongoCollection *mongo.Collection
}

func (r *CardRepo) InsertCard(card *types.Card) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne((context.Background()), card)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *CardRepo) DeleteCard(cardID string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(cardID)
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

func (r *CardRepo) FindCardByID(cardID string) (*types.Card, error) {
	id, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", id}}

	var card types.Card

	err = r.MongoCollection.FindOne(context.Background(), filter).Decode(&card)
	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &card, nil
}

func (r *CardRepo) FindAllCards() ([]types.Card, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var cards []types.Card

	err = results.All(context.Background(), &cards)
	if err != nil {
		return nil, fmt.Errorf("Find all uses results decode error %s", err.Error())
	}

	return cards, nil
}

func (r *CardRepo) UpdateName(cardID string, newName string) error {
	id, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", newName}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *CardRepo) AddTransaction(cardID string, spend *types.Spend) error {
	id, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$push", bson.D{{"transactions", spend}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to add transaction to card: %w", err)
	}

	return nil
}

func (r *UserRepo) RemoveTransaction(cardID string, transactionID string) error {
	cardObjID, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return err
	}

	transactionObjID, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", cardObjID}}
	update := bson.D{{"$pull", bson.D{{"transactions", bson.D{{"_id", transactionObjID}}}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to remove transaction from card: %w", err)
	}

	return nil
}
