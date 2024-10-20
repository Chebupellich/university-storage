package routes

import (
	"net/http"
	"server/controllers"
	"server/middleware"

	"github.com/gorilla/mux"
)

func SetUserRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/register", controllers.RegisterHandler).Methods(http.MethodPost)

	userRoutes := r.PathPrefix("/users").Subrouter()
	userRoutes.Use(middleware.AuthMiddleware)

	userRoutes.HandleFunc("", controllers.GetUsers).Methods(http.MethodGet)
	userRoutes.HandleFunc("/{name}", controllers.GetUser).Methods(http.MethodGet)
	userRoutes.HandleFunc("/admin-access", controllers.AdminAccessHandler).Methods(http.MethodPost)

	// Admin routes allow CRUD operations with users
	adminRoutes := userRoutes.PathPrefix("").Subrouter()
	adminRoutes.Use(middleware.RoleMiddleware("admin"))
	adminRoutes.HandleFunc("/", controllers.CreateUser).Methods(http.MethodPost)
	adminRoutes.HandleFunc("/{name}", controllers.UpdateUser).Methods(http.MethodPut)
	adminRoutes.HandleFunc("/{name}", controllers.DeleteUser).Methods(http.MethodDelete)

}
