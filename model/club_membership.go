package model

type ClubMembership struct {
	ClubID string `bson:"club_id"`
	UserID string `bson:"user_id"`
}
