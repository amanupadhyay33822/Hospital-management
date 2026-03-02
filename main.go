package main

import (
	"go-auth/internal/config"
	"go-auth/internal/logger"
	"go-auth/internal/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var userCollection *mongo.Collection

func main() {
	config.ConnectDB()
	var router *mux.Router
	router = mux.NewRouter()
	router.Use(logger.Logger)
	routes.UserRoutes(router)
	routes.DoctorRoutes(router)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

