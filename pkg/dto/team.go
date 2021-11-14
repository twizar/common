package dto

type Team struct {
	ID     string  `json:"id" bson:"_id"`
	Name   string  `json:"name" bson:"name"`
	Rating float64 `json:"rating" bson:"rating"`
	League string  `json:"league" bson:"league"`
}
