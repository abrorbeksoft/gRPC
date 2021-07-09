package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
}

func (user *User) SetName(name string)  {
	user.Name=name
}

func (user *User) GetName() string  {
	return user.Name
}

func main()  {
	fmt.Println("Hello world")
	user:=User{};
	user.SetName("Abrorbek")

	fmt.Println(user.GetName())
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Server is started")

	})

	err:=http.ListenAndServe(":8081",nil)
	if err !=nil{
		fmt.Println(err)
	}
}