package service

import (
	"my_project/dao"
	"my_project/model"
)

func AddStudent(student *model.Student) error {
	return dao.CreateStudent(student)
}

func ListStudents() ([]model.Student, error) {
	return dao.GetAllStudents()
}
func GetStudentByID(id string) (*model.Student, error) {
	return dao.GetStudentByID(id)
}

func UpdateStudent(id string, student *model.Student) error {
	return dao.UpdateStudent(id, student)
}

func DeleteStudent(id string) error {
	return dao.DeleteStudent(id)
}
