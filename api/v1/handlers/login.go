package handlers

type User struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age int `json:"age"`
}

func Login(name, surname string) User {
	return User{Name: name,Surname: surname,Age: 28}
}
