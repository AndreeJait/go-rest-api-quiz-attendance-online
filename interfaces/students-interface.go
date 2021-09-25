package interfaces

import (
	"time"

	"github.com/AndreeJait/go-rest-api-quiz-attendance-online/entity"
	"github.com/AndreeJait/go-rest-api-quiz-attendance-online/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentsInterface interface {
	FindAll(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type studentsController struct {
	service service.StudentService
}

func New(service service.StudentService) StudentsInterface {
	return &studentsController{
		service: service,
	}
}

func (c *studentsController) FindAll(ctx *gin.Context) {
	ctx.JSON(201, c.service.FindAll())
}

func (c *studentsController) Register(ctx *gin.Context) {
	var student entity.Students
	var studentRequest entity.StudentsRequest
	ctx.BindJSON(&studentRequest)

	t, err := time.Parse("2006-01-02", studentRequest.BirthDate)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
	} else {
		student.ID = primitive.NewObjectID()
		student.BirthDate = t
		student.BirthPlace = studentRequest.BirthPlace
		student.Name = studentRequest.Name
		student.ClassCode = studentRequest.ClassCode
		student.Email = studentRequest.Email
		student.Gender = studentRequest.Gender
		student.Password = studentRequest.Password
		student.Address = studentRequest.Address
		result, err := c.service.Register(student)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Internal Server Error",
			})
		} else {
			ctx.JSON(201, gin.H{
				"msg":    "Succes to insert data.",
				"result": result,
			})
		}
	}

}
