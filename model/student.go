package model

type Student struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100)"`
	Age   int
	Email string `gorm:"type:varchar(100);unique"`
}
