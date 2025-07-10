package controller

import (
	"my_project/model"
	"my_project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStudent(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "添加成功", "student": student})
}

func GetStudents(c *gin.Context) {
	students, err := service.ListStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"students": students})
}

// 获取单个学生信息
func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := service.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "学生不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"student": student})
}

// 更新学生信息
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateStudent(id, &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "student": student})
}

// 删除学生信息
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteStudent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
