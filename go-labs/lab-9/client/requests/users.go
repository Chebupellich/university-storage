package requests

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age" json:"age"`
}

func CreateUser() {

}

func GetUsers() {

}

func GetUser() {

}

func UpdateUser() {

}

func DeleteUser() {

}
