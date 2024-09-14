package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (self *repository) createClubIndexes(ctx context.Context) error {
	collection := self.mongoDB.Collection(CLUB_COLLECTION_NAME)

	// Define a unique index for the "name" field
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}}, // 1 means ascending order
		Options: options.Index().SetUnique(true),
	}

	// Create the index
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	log.Println("Unique index on 'name' field created successfully")
	return nil
}

func (self *repository) createClubMembershipIndexes(ctx context.Context) error {
	collection := self.mongoDB.Collection(CLUB_MEMBERSHIP_COLLECTION_NAME)

	// Define a unique compound index for the combination of "club_id" and "user_id"
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"club_id", 1}, {"user_id", 1}}, // Compound index on club_id and user_id
		Options: options.Index().SetUnique(true),
	}

	// Create the index
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	log.Println("Unique compound index on 'club_id' and 'user_id' created successfully")
	return nil
}
