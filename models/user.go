package models

type User struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
	Phone string `bson:"phone"`
}
