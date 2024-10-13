package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

var httpClient http.Client = http.Client{}

func main() {
	req, err := http.NewRequest()
}

func UI(req *http.Request) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		for {
			fmt.Print("> ")
			if scanner.Scan() {

			}
		}
	}
}

func SendRequest(r *http.Request) *http.Response {
	resp, err := httpClient.Do(r)
	if err != nil {
		SendRequestError(err)
		return nil
	}
	defer resp.Body.Close()
	return resp
}
