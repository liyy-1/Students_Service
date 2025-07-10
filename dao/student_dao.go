package dao

import (
	"my_project/config"
	"my_project/model"
)

func CreateStudent(student *model.Student) error {
	return config.DB.Create(student).Error
}

func GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	err := config.DB.Find(&students).Error
	return students, err
}

// 根据 ID 获取单个学生
func GetStudentByID(id string) (*model.Student, error) {
	var student model.Student
	err := config.DB.First(&student, id).Error
	return &student, err
}

// 更新学生信息
func UpdateStudent(id string, data *model.Student) error {
	var student model.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		return err
	}

	student.Name = data.Name
	student.Age = data.Age
	student.Email = data.Email

	return config.DB.Save(&student).Error
}

// 删除学生信息
func DeleteStudent(id string) error {
	return config.DB.Delete(&model.Student{}, id).Error
}
