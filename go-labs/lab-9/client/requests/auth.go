package requests

import (
	"bytes"
	"client/exceptions"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Logout() {
	IsAuth = false
	accessToken = ""
	rootToken = ""
}

func Login(name string, password string) bool {
	body, _ := json.Marshal(map[string]string{"login": name, "password": password})

	client := http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		exceptions.LoginError(err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		IsAuth = true

		var b []byte
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			exceptions.ParseError(err)
			return false
		}

		if SetTokens(b) {
			return true
		}
	}

	return false
}

func Register(name string, password string) bool {
	body, _ := json.Marshal(map[string]string{"login": name, "password": password})

	client := http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		exceptions.RegistrationError(err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var b []byte
		b, err = io.ReadAll(resp.Body)
		if err != nil {
			exceptions.ParseError(err)
			return false
		}

		if SetTokens(b) {
			return true
		}
	}

	return true
}

func CheckAuth() bool {
	if len(accessToken) == 0 {
		return false
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/auth/%s", baseURL, accessToken), nil)

	resp, err := client.Do(req)

	if err == nil && resp.StatusCode == http.StatusUnauthorized {
		if RefreshAccessToken() {
			req, _ := http.NewRequest("POST", fmt.Sprintf("%s/auth/%s", baseURL, accessToken), nil)
			resp, err := client.Do(req)

			if err == nil && resp.StatusCode == http.StatusAccepted {
				IsAuth = true
				return true
			}
		} else {
			return false
		}
	}

	if err == nil && resp.StatusCode == http.StatusAccepted {
		IsAuth = true
		return true
	}

	return false
}

func RefreshAccessToken() bool {
	if len(refreshToken) == 0 {
		return false
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/refresh/%s", baseURL, refreshToken), nil)
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode == http.StatusNotAcceptable {
		return false
	}

	defer resp.Body.Close()

	var b []byte
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		exceptions.ParseError(err)
		return false
	}

	if SetTokens(b) {
		return true
	}

	return false
}

func SetTokens(data []byte) bool {
	var tokens TokenResponse
	err := json.Unmarshal(data, &tokens)
	if err != nil {
		return false
	}

	accessToken = tokens.Access

	return true
}
