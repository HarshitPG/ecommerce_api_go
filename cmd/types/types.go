package types

import (
	"time"
)
type UserStore interface{
	GetUserByEmail(email string) (*User , error)
	GetuserByID(id int)(*User,error)
	CreateUser(User)error
}

// type mockUserStore struct{

// }
// func GetUserByEmail(email string)(*User, error){
// 	return nil,nil
// }
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}