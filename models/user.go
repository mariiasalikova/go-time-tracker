package models

type User struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	PassportNumber string `gorm:"unique" json:"passportNumber"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}
