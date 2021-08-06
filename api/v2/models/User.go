package models

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age string `json:"age"`
	Salary string `json:"salary"`
}
