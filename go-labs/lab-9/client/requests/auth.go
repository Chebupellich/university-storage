package requests

import (
	"fmt"
	"net/http"

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

func CheckAuth() {

}
