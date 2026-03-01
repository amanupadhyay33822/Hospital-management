package routes

import (
	"go-auth/internal/handlers"
	"go-auth/internal/utils"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/upload", utils.UploadImage).Methods("POST")
}
func DoctorRoutes(router *mux.Router) {
	router.HandleFunc("/doctors/create", handlers.CreateDoctor).Methods("POST")
	router.HandleFunc("/doctors", handlers.GetAllDoctors).Methods("GET")
	router.HandleFunc("/doctors/{id}", handlers.GetDoctorByID).Methods("GET")
	router.HandleFunc("/doctors/{id}", handlers.UpdateDoctor).Methods("PUT")
	router.HandleFunc("/doctors/{id}", handlers.DeleteDoctor).Methods("DELETE")
}
