package repositories

import (
	"context"
	"fmt"
	"money-minder/internal/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	MongoCollection *mongo.Collection
}

func (r *UserRepo) InsertUser(usr *types.User) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne((context.Background()), usr)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepo) DeleteUser(usr *types.User) (interface{}, error) {
	result, err := r.MongoCollection.DeleteOne((context.Background()), usr)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepo) FindUserByID(usrID string) (*types.User, error) {
	id, err := primitive.ObjectIDFromHex(usrID)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", id}}

	var usr types.User

	err = r.MongoCollection.FindOne(context.Background(), filter).Decode(&usr)
	if err != nil {

		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &usr, nil
}

func (r *UserRepo) FindAllUsers() ([]types.User, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var usrs []types.User

	err = results.All(context.Background(), &usrs)
	if err != nil {
		return nil, fmt.Errorf("Find all uses results decode error %s", err.Error())
	}

	return usrs, nil
}

func (r *UserRepo) UpdateName(usrID string, newName string) error {
	id, err := primitive.ObjectIDFromHex(usrID)
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

func (r *UserRepo) UpdatePassword(usrID string, newPassword string) error {
	id, err := primitive.ObjectIDFromHex(usrID)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"password", newPassword}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) AddCard(userID string, cardID string, cardRepo *CardRepo) error {

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	card, err := cardRepo.FindCardByID(cardID)
	if err != nil {
		return fmt.Errorf("failed to find card: %w", err)
	}
	if card == nil {
		return fmt.Errorf("card not found")
	}

	filter := bson.D{{"_id", userObjectID}}
	update := bson.D{{"$push", bson.D{{"cards", card}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to add card to user: %w", err)
	}

	return nil
}

func (r *UserRepo) RemoveCard(userID string, cardID string, cardRepo *CardRepo) error {

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	card, err := cardRepo.FindCardByID(cardID)
	if err != nil {
		return fmt.Errorf("failed to find card: %w", err)
	}
	if card == nil {
		return fmt.Errorf("card not found")
	}

	filter := bson.D{{"_id", userObjectID}}
	update := bson.D{{"$pull", bson.D{{"cards", bson.D{{"$eq", card}}}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete card from user: %w", err)
	}

	return nil
}

func (r *UserRepo) AddSpend(userID string, spendID string, spendRepo *SpendRepo) error {

	userObjID, err := primitive.ObjectIDFromHex(userID)
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

	filter := bson.D{{"_id", userObjID}}
	update := bson.D{{"$push", bson.D{{"spends", spend}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to add spend to user: %w", err)
	}

	return nil
}

func (r *UserRepo) RemoveSpend(userID string, spendID string, spendRepo *SpendRepo) error {

	userObjectID, err := primitive.ObjectIDFromHex(userID)
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

	filter := bson.D{{"_id", userObjectID}}
	update := bson.D{{"$pull", bson.D{{"spends", bson.D{{"$eq", spend}}}}}}

	_, err = r.MongoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to delete spend from user: %w", err)
	}

	return nil
}
