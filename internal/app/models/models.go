package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BsonProduct struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name, omitempty"`
	Info  string             `bson:"info, omitempty"`
	Img   string             `bson:"img, omitempty"`
	Price string             `bson:"price, omitempty"`
}
type BsonProductList []BsonProduct

type HandlerError struct {
	Error string `json:"error"`
}

type HandlerStatus struct {
	Status string `json:"status"`
}
