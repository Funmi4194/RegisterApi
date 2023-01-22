package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//created a struct that satisfies both mus package and mongo database
//note bson can be removed then mongo with convert the struct fields to get its name which the frontend guys can work with
//omitempty mean it can be ommited
type User struct {
	ID         primitive.ObjectID `json:"id" bson:"ID,omitempty"`
	Name       string             `json:"name,omitempty" bson:"Name,omitempty" validate:"required"`
	Age        int                `json:"age,omitempty" bson:"Age,omitempty" validate:"required"`
	Profession string             `json:"profession,omitempty" bson:"Job,omitempty" validate:"required"`
	Location   string             `json:"Location,omitempty" bson:"Landmark,omitempty" validate:"required"`
}
