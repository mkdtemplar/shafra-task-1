package models

type User struct {
	ID          int64  `gorm:"type:smallint;primaryKey" json:"id" binding:"required"`
	NameSurname string `gorm:"type:text" json:"name_surname" binding:"required"`
	Age         int64  `gorm:"type:smallint" json:"age" binding:"required"`
}
