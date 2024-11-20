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

func (r *CardRepo) AddSpend(cardId string, spendId string, spendRepo *SpendRepo) error {

	id, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return err
	}

	spend, err := spendRepo.FindSpendByID(spendId)

	if err != nil {
		return fmt.Errorf("failed to find card: %w", err)
	}
	if spend == nil {
		return fmt.Errorf("card not found")
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$push", bson.D{{"spends", spend}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to add spends to card: %w", err)
	}

	return nil
}

func (r *CardRepo) RemoveSpend(cardID string, spendID string, spendRepo *SpendRepo) error {

	cardObjectId, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	spend, err := spendRepo.FindSpendByID(spendID)
	if err != nil {
		return fmt.Errorf("failed to find spend: %w", err)
	}
	if spend == nil {
		return fmt.Errorf("spend not found")
	}

	filter := bson.D{{"_id", cardObjectId}}
	update := bson.D{{"$pull", bson.D{{"spends", bson.D{{"$eq", spend}}}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete spend from card: %w", err)
	}

	return nil
}

func (r *CardRepo) GetCardsByUserID(userID string) ([]*types.Card, error) {

	ownerID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid userID: %w", err)
	}

	filter := bson.M{"owner": ownerID}
	var cards []*types.Card

	cursor, err := r.MongoCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find cards: %w", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var card types.Card
		if err := cursor.Decode(&card); err != nil {
			return nil, fmt.Errorf("failed to decode card: %w", err)
		}
		cards = append(cards, &card)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return cards, nil
}
