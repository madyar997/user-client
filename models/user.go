package models

type User struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int32  `json:"age"`
	Password string `json:"password"`
}
