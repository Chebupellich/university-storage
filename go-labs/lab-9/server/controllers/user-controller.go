package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"server/handlers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name" json:"name"`
	Age  int                `bson:"age" json:"age"`
}

var userCollection *mongo.Collection

func InitUserController(collection *mongo.Database) {
	userCollection = collection.Collection("users")
}

func ValidateInput(user User) bool {
	if reflect.TypeOf(user.Name).String() == "string" ||
		user.Name != "" ||
		reflect.TypeOf(user.Age).String() == "int" ||
		(user.Age > 0 && user.Age < 150) {

		return true
	}
	return false
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GETUSERS")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	users := []User{}

	cursor, err := userCollection.Find(ctx, filter)

	if err != nil {
		fmt.Println("ARR: ", err)
		handlers.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var usr User
		if err := cursor.Decode(&usr); err != nil {
			fmt.Println("ARR: ", err)
			handlers.HandleError(w, err, http.StatusInternalServerError)
			return
		}
		users = append(users, usr)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var usr User
	name := mux.Vars(r)["name"]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"name": name}).Decode(&usr)
	if err != nil {
		handlers.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var usr User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		handlers.HandleError(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("CREATE: ", string(body))
	err = json.Unmarshal(body, &usr)
	if err != nil {
		handlers.HandleError(w, err, http.StatusBadRequest)
		return
	}

	if !ValidateInput(usr) {
		handlers.HandleError(w, fmt.Errorf("incorrect value"), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	usr.Id = primitive.NewObjectID()
	_, err = userCollection.InsertOne(ctx, usr)
	if err != nil {
		handlers.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usr)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	var usr User
	upd := bson.M{"$set": bson.M{}}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handlers.HandleError(w, err, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &usr)
	if err != nil {
		handlers.HandleError(w, err, http.StatusBadRequest)
		return
	}

	if usr.Name != "" {
		upd["$set"].(bson.M)["name"] = usr.Name
	}
	if usr.Age > 0 && usr.Age < 150 {
		upd["$set"].(bson.M)["age"] = usr.Age
	}

	if len(upd["$set"].(bson.M)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = userCollection.UpdateOne(ctx, bson.M{"name": name}, upd)
	if err != nil {
		handlers.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User{Id: usr.Id, Name: usr.Name, Age: usr.Age})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := userCollection.DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		handlers.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
