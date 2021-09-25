package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Students struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=10,max=255"`
	ClassCode  []string           `json:"class_code"`
	Address    string             `json:"address"`
	Gender     *string            `json:"gender" validate:"required,min=1,max=2"`
	Email      *string            `json:"email" validate:"required"`
	Password   *string            `json:"password" validate:"required"`
	BirthDate  time.Time          `json:"birth_date"`
	BirthPlace *string            `json:"birth_place" validate:"required"`
	Role       int                `json:"role"`
}

type StudentsRequest struct {
	Name       *string  `json:"name" validate:"required,min=10,max=255"`
	ClassCode  []string `json:"class_code"`
	Address    string   `json:"address"`
	Gender     *string  `json:"gender" validate:"required,min=1,max=2"`
	Email      *string  `json:"email" validate:"required"`
	Password   *string  `json:"password" validate:"required"`
	BirthDate  string   `json:"birth_date"`
	BirthPlace *string  `json:"birth_place" validate:"required"`
	Role       int      `json:"role"`
}
