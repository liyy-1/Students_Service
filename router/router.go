package router

import (
	"my_project/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/students", controller.AddStudent)
	r.GET("/students", controller.GetStudents)
	r.GET("/students/:id", controller.GetStudentByID)
	r.PUT("/students/:id", controller.UpdateStudent)
	r.DELETE("/students/:id", controller.DeleteStudent)

	return r
}
