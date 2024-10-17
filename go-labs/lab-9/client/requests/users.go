package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"client/exceptions"
)

func SendRequest(req http.Request) {
	req.Header.Set("Authorization", "Bearer "+accessToken)

}

func CreateUser(name string, age int) bool {
	body, _ := json.Marshal(SendUser{Name: name, Age: age})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users", baseURL), bytes.NewBuffer(body))
	if err != nil {
		exceptions.CreateRequestError(err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	// TODO: make in header access token

	resp, err := client.Do(req)
	if err != nil {
		exceptions.GeneralError(err)
		return false
	}

	if status := resp.StatusCode; status != http.StatusCreated {
		return false
	}

	return true
}

func GetUsers() []User {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users?page=1&limit=10", baseURL), nil)
	if err != nil {
		exceptions.CreateRequestError(err)
		return nil
	}

	// TODO: make in header access token

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode == 400 {
		exceptions.GeneralError(err)
		return nil
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		exceptions.ParseError(err)
		return nil
	}

	var users []User
	err = json.Unmarshal(b, &users)
	if err != nil {
		exceptions.GetRequestError(err)
		return nil
	}

	return users
}

func GetUser(name string) *User {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", baseURL, name), nil)
	if err != nil {
		exceptions.CreateRequestError(err)
		return nil
	}

	// TODO: make in header access token

	resp, err := client.Do(req)
	if err != nil {
		exceptions.GeneralError(err)
		return nil
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		exceptions.ParseError(err)
		return nil
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		exceptions.GetRequestError(err)
		return nil
	}

	return &user
}

func UpdateUser(oldName string, newName string, age int) *User {
	body, _ := json.Marshal(SendUser{Name: newName, Age: age})

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/%s", baseURL, oldName), bytes.NewBuffer(body))
	if err != nil {
		exceptions.CreateRequestError(err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	// TODO: make in header access token

	resp, err := client.Do(req)
	if err != nil {
		exceptions.GeneralError(err)
		return nil
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		exceptions.ParseError(err)
		return nil
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		exceptions.GetRequestError(err)
		return nil
	}

	return &user
}

func DeleteUser(name string) bool {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/%s", baseURL, name), nil)
	if err != nil {
		exceptions.CreateRequestError(err)
		return false
	}

	// TODO: make in header access token

	resp, err := client.Do(req)
	if err != nil {
		exceptions.GeneralError(err)
		return false
	}
	if status := resp.StatusCode; status != http.StatusNoContent {
		return false
	}
	return true
}
