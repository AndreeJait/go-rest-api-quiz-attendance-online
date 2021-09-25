package main

import (
	"errors"
	"log"
	"os"

	"github.com/AndreeJait/go-rest-api-quiz-attendance-online/interfaces"
	"github.com/AndreeJait/go-rest-api-quiz-attendance-online/service"
	"github.com/gin-gonic/gin"
)

var (
	studentsService    service.StudentService       = service.NewStudentsService()
	studentsController interfaces.StudentsInterface = interfaces.NewStudentsInterface(studentsService)
)

func main() {
	server := gin.New()
	server.Use(gin.Logger())

	err := errors.New("Error")

	for err != nil {
		err = service.ConnectDB()
		if err != nil {
			log.Fatal(err)
		}
	}

	server.POST("/add", studentsController.RegisterStudents)

	// server.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"code":    200,
	// 		"message": "Success connect to Gin",
	// 	})
	// })

	server.Run(":" + os.Getenv("PORT"))
}
