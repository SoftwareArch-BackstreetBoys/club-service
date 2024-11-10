package model

type Club struct {
	ID                string `bson:"_id" json:"id"`
	Name              string `bson:"name" json:"name"`
	Description       string `bson:"description" json:"description"`
	CreatedByID       string `bson:"created_by_id" json:"created_by_id"`
	CreatedByFullName string `bson:"created_by_full_name" json:"created_by_full_name"`
}
