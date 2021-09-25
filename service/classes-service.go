package service

import (
	"github.com/AndreeJait/go-rest-api-quiz-attendance-online/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassesService interface {
	FindAll() []entity.Claseses
	Register(student entity.Claseses) (*mongo.InsertOneResult, error)
}

type classesService struct {
}

func NewClasses() ClassesService {
	return &classesService{}
}
