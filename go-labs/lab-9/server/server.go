package main

import (
	"fmt"
	"net/http"
	"server/controllers"
	"server/dbconnect"
	"server/routes"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Database

func main() {
	userCollection = dbconnect.DBConnect("mongodb://localhost:27017/", "sacred_base")
	controllers.InitUserController(userCollection)

	r := mux.NewRouter()
	routes.SetUserRoutes(r)

	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", r)
}
