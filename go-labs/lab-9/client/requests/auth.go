package requests

import (
	"client/exceptions"
	"fmt"
	"net/http"
	"os"

	tm "github.com/buger/goterm"
)

func Login(nick string, password string) (*http.Response, error) {
	fmt.Println(nick, password)

	client := http.Client{}
	resp, err := client.Get("8.8.8.8") // TODO: Post req

	if err != nil {
		fmt.Println(tm.Color(err.Error(), tm.RED))
	}

	return resp, err
}

func Register(nick string, password string) (*http.Response, error) {
	client := http.Client{}
	resp, err := client.Get("8.8.8.8") // TODO: Post req

	if err != nil {
		fmt.Println(tm.Color(err.Error(), tm.RED))
	}

	return resp, err
}

func CheckAuth() (bool, string) {
	token, err := os.ReadFile("token.txt")
	if err != nil {
		exceptions.CannotReadFileError(err)
	}

	if len(token) == 0 {
		return false, ""
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/%s", baseURL, string(token)), nil)

	resp, err := client.Do(req)

	if err != nil {
		exceptions.GeneralError(err)
		return false, string(token)
	}
	if status := resp.StatusCode; status != http.StatusAccepted {
		return false, string(token)
	}

	err = os.WriteFile("output.txt", token, 0777)
	if err != nil {
		exceptions.CannotReadFileError(err)
	}

	return true, string(token)
}

func RefreshAccessToken() {
	// Get refresh token from file
	// Send request to refresh access token
	// Get new access token
	// Set new access token
	// if success return true
	// else drop gorutines and load login page
}
