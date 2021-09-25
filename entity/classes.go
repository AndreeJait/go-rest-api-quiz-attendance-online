package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Claseses struct {
	ID          primitive.ObjectID `bson:"_id"`
	Code        *string            `json:"code" validate:"required"`
	Name        *string            `json:"name" validate:"required,max=255"`
	Owner       primitive.ObjectID `bson:"owner"`
	Description string             `json:"description"`
}

type ClasesesRequest struct {
	Code        *string            `json:"code" validate:"required"`
	Name        *string            `json:"name" validate:"required,max=255"`
	Owner       primitive.ObjectID `bson:"owner"`
	Description string             `json:"description"`
}
