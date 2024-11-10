package repository

import (
	"context"
	"errors"

	"github.com/SoftwareArch-BackstreetBoys/club-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateClub inserts a new club into the clubs collection.
func (self *repository) CreateClub(ctx context.Context, club model.Club) (*model.Club, error) {
	collection := self.mongoDB.Collection(CLUB_COLLECTION_NAME)

	// Insert the club into the MongoDB collection
	_, err := collection.InsertOne(ctx, club)
	if err != nil {
		// Check if the error is due to a duplicate key (i.e., club name already exists)
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.New("club name already exists")
		}
		return nil, err
	}

	return &club, nil
}

// GetAllClubs retrieves all clubs from the clubs collection.
func (self *repository) GetAllClubs(ctx context.Context) ([]model.Club, error) {
	collection := self.mongoDB.Collection(CLUB_COLLECTION_NAME)

	// Find all documents in the collection
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	clubs := make([]model.Club, 0)
	if err = cursor.All(ctx, &clubs); err != nil {
		return nil, err
	}

	return clubs, nil
}

// GetClub retrieves a single club by its ID from the clubs collection.
func (self *repository) GetClub(ctx context.Context, clubID string) (*model.Club, error) {
	collection := self.mongoDB.Collection(CLUB_COLLECTION_NAME)

	// Find the club by its _id field
	var club model.Club
	err := collection.FindOne(ctx, bson.D{{"_id", clubID}}).Decode(&club)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("club not found")
		}
		return nil, err
	}

	return &club, nil
}

// GetClubs retrieves multiple clubs by their IDs.
func (self *repository) GetClubs(ctx context.Context, clubIDs []string) ([]model.Club, error) {
	collection := self.mongoDB.Collection(CLUB_COLLECTION_NAME)

	// Find clubs whose _id field is in the provided list of clubIDs
	filter := bson.D{{"_id", bson.D{{"$in", clubIDs}}}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	clubs := make([]model.Club, 0)
	if err = cursor.All(ctx, &clubs); err != nil {
		return nil, err
	}

	return clubs, nil
}

// SearchClubs searches for clubs by a keyword in the name or description fields.
func (self *repository) SearchClubs(ctx context.Context, keyword string) ([]model.Club, error) {
	collection := self.mongoDB.Collection(CLUB_COLLECTION_NAME)

	// Search for clubs where name or description contains the keyword
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"name", bson.D{{"$regex", keyword}, {"$options", "i"}}}},
			bson.D{{"description", bson.D{{"$regex", keyword}, {"$options", "i"}}}},
		}},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	clubs := make([]model.Club, 0)
	if err = cursor.All(ctx, &clubs); err != nil {
		return nil, err
	}

	return clubs, nil
}
