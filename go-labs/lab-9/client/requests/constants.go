package requests

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
	Age  int                `json:"age"`
}

type SendUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type RootAccess struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

var client = http.Client{
	Timeout: 10 * time.Second,
}

const baseURL = "http://localhost:8080/users"

type TokenResponse struct {
	Access string `json:"Access"`
}

var accessToken string
var refreshToken string

var IsAuth bool = false

var rootToken string
