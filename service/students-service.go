package service

import (
	"context"
	"time"

	"github.com/AndreeJait/go-rest-api-quiz-attendance-online/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentService interface {
	FindAllStudents() []entity.Students
	RegisterStudents(student entity.Students) (*mongo.InsertOneResult, error)
}

type studentService struct {
	students []entity.Students
}

func NewStudentsService() StudentService {
	return &studentService{}
}

func (s *studentService) FindAllStudents() []entity.Students {
	return s.students
}

func (s *studentService) RegisterStudents(student entity.Students) (*mongo.InsertOneResult, error) {
	studentsCollection := MI.DB.Collection("students")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := studentsCollection.InsertOne(ctx, student)

	if err != nil {
		return nil, err
	}

	return result, nil
}
