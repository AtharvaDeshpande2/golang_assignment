package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Items struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty"`
	Price    int                `json:"price,omitempty"`
	Quantity int                `json:"quantity,omitempty"`
}
type RegisterUser struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"`
	Age     int                `json:"age,omitempty"`
	Address string             `json:"add,omitempty"`
}
