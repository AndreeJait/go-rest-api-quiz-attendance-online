package service

import (
	"context"
	"time"

	"github.com/AndreeJait/go-rest-api-quiz-attendance-online/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClassesService interface {
	FindAllClass() []entity.Claseses
	AddClass(claseses entity.Claseses) (*mongo.InsertOneResult, error)
}

type classesService struct {
	classes []entity.Claseses
}

func NewClassesService() ClassesService {
	return &classesService{}
}

func (c *classesService) FindAllClass() []entity.Claseses {
	return c.classes
}
func (c *classesService) AddClass(claseses entity.Claseses) (*mongo.InsertOneResult, error) {
	clasesCollection := MI.DB.Collection("classes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := clasesCollection.InsertOne(ctx, claseses)

	if err != nil {
		return nil, err
	}

	return result, nil
}
