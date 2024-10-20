package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"client/exceptions"
)

func SendRequest(req *http.Request) *http.Response {
	if rootToken != "" {
		req.Header.Set("Authorization", "Bearer "+rootToken)
	} else {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}

	resp, err := client.Do(req)

	if err != nil {
		exceptions.GeneralError(err)
	} else if resp.StatusCode == http.StatusUnauthorized {
		IsAuth = false
		return nil
	}

	return resp
}

func CreateUser(name string, age int) bool {
	var usr SendUser
	usr.Name = name
	usr.Age = age

	body, _ := json.Marshal(usr)
	fmt.Println("CREATE: ", string(body), age)
	req, err := http.NewRequest("POST", "http://localhost:8080/users/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		exceptions.CreateRequestError(err)
		return false
	}

	resp := SendRequest(req)
	if resp == nil {
		return false
	}

	if status := resp.StatusCode; status != http.StatusCreated {
		return false
	}

	return true
}

func GetUsers() []User {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		exceptions.CreateRequestError(err)
		return nil
	}

	resp := SendRequest(req)
	if resp == nil || resp.StatusCode == 400 {
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseURL, name), nil)
	if err != nil {
		exceptions.CreateRequestError(err)
		return nil
	}

	resp := SendRequest(req)
	if resp == nil {
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

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", baseURL, oldName), bytes.NewBuffer(body))
	if err != nil {
		exceptions.CreateRequestError(err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	resp := SendRequest(req)
	if resp == nil {
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseURL, name), nil)
	if err != nil {
		exceptions.CreateRequestError(err)
		return false
	}

	resp := SendRequest(req)
	if resp == nil {
		return false
	}

	if status := resp.StatusCode; status != http.StatusNoContent {
		return false
	}
	return true
}

func GetAdminAccess(passwd string) bool {
	var input struct {
		Secret string `json:"secret"`
	}
	input.Secret = passwd

	body, _ := json.Marshal(input)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin-access", baseURL), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		exceptions.CreateRequestError(err)
		return false
	}

	resp := SendRequest(req)
	if resp == nil {
		return false
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		exceptions.ParseError(err)
		return false
	}

	var root RootAccess
	err = json.Unmarshal(b, &root)
	if err != nil {
		exceptions.GetRequestError(err)
		return false
	}

	fmt.Println("\033[32m * GET ROOT ACCESS * \033[0m")
	rootToken = root.Token

	return true
}
