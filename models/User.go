package models

type User struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age int32 `json:"age"`
}

