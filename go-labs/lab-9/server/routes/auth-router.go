package routes

import (
	"net/http"
	"server/controllers"

	"github.com/gorilla/mux"
)

func SetAuthRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/register", controllers.RegisterHandler).Methods(http.MethodPost)
}
