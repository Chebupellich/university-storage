package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	getRoot()
	getHello()
	postData()
}

func getRoot() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println("Ошибка при GET-запросе на /:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ответ на GET /:", string(body))
}

func getHello() {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		fmt.Println("Ошибка при GET-запросе на /hello:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ответ на GET /hello:", string(body))
}

func postData() {
	data := map[string]interface{}{
		"key":    "value",
		"number": 123,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка при кодировании JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/data", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при POST-запросе на /data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ответ на POST /data:", string(body))
}
