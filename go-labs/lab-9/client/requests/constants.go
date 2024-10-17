package requests

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age" json:"age"`
}

type SendUser struct {
	Name string
	Age  int
}

var client = http.Client{
	Timeout: 10 * time.Second,
}

const baseURL = "http://localhost:8080"

var accessToken string

//var refreshToken string
