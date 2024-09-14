package model

type Club struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
}
