package repository

import (
	"context"
	"errors"

	"github.com/SoftwareArch-BackstreetBoys/club-service/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateClubMemberShip adds a membership record for a user in a club.
func (self *repository) CreateClubMemberShip(ctx context.Context, clubID string, userID string) error {
	clubCollection := self.mongoDB.Collection(CLUB_COLLECTION_NAME)
	membershipCollection := self.mongoDB.Collection(CLUB_MEMBERSHIP_COLLECTION_NAME)

	// Check if the club exists
	var club model.Club
	err := clubCollection.FindOne(ctx, bson.D{{"_id", clubID}}).Decode(&club)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("club not found")
		}
		return err
	}

	// Create a new ClubMembership instance
	membership := model.ClubMembership{
		ClubID: clubID,
		UserID: userID,
	}

	// Insert the membership into the collection
	_, err = membershipCollection.InsertOne(ctx, membership)
	if err != nil {
		// Check if the error is due to a duplicate key (i.e., user already joined the club)
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("user is already a member of the club")
		}
		return err
	}

	return nil
}

// GetClubMemberShip retrieves the club membership for a user in a specific club.
func (self *repository) GetClubMemberShip(ctx context.Context, clubID string, userID string) (*model.ClubMembership, error) {
	collection := self.mongoDB.Collection(CLUB_MEMBERSHIP_COLLECTION_NAME)

	// Filter to find the membership based on club_id and user_id
	filter := bson.D{
		{"club_id", clubID},
		{"user_id", userID},
	}

	var membership model.ClubMembership
	err := collection.FindOne(ctx, filter).Decode(&membership)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("membership not found")
		}
		return nil, err
	}

	return &membership, nil
}

// DeleteClubMemberShip removes a membership record for a user in a club.
func (self *repository) DeleteClubMemberShip(ctx context.Context, clubID string, userID string) error {
	collection := self.mongoDB.Collection(CLUB_MEMBERSHIP_COLLECTION_NAME)

	// Filter to find the membership based on club_id and user_id
	filter := bson.D{
		{"club_id", clubID},
		{"user_id", userID},
	}

	// Delete the membership
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if any documents were deleted
	if result.DeletedCount == 0 {
		return errors.New("membership not found")
	}

	return nil
}

// GetJoinedClubIDS retrieves the list of club IDs the user has joined.
func (self *repository) GetJoinedClubIDS(ctx context.Context, userID string) ([]string, error) {
	collection := self.mongoDB.Collection(CLUB_MEMBERSHIP_COLLECTION_NAME)

	// Filter to find memberships by user_id
	filter := bson.D{
		{"user_id", userID},
	}

	// Find all matching memberships
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var clubIDs []string
	for cursor.Next(ctx) {
		var membership model.ClubMembership
		if err := cursor.Decode(&membership); err != nil {
			return nil, err
		}
		clubIDs = append(clubIDs, membership.ClubID)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return clubIDs, nil
}
