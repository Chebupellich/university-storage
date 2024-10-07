package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func handleDefaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "default route")
}

func handleHelloRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, on hello route")
}

func handleDataRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		fmt.Println("Полученный JSON:", string(body))

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "JSON получен и выведен на консоль.")
	} else {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Printf("Request: %s %s in %v\n", r.Method, r.URL, time.Now().Format("2006.01.02 15:04:05"))
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleDefaultRoute)
	mux.HandleFunc("/hello", handleHelloRoute)
	mux.HandleFunc("/data", handleDataRoute)

	loggedMux := loggingMiddleware(mux)

	fmt.Println("Сервер запущен на порту 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
