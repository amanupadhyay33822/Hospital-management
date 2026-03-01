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

// -------------------- REGISTER --------------------
// func Register(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)

// 	// Check if user exists
// 	filter := bson.M{"email": user.Email}
// 	var existingUser User

// 	err := userCollection.FindOne(context.Background(), filter).Decode(&existingUser)

// 	fmt.Printf("Existing user: %+v\n", existingUser)
// 	if err == nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]string{
// 			"error": "User already exists",
// 		})
// 		return
// 	}

// 	// Hash password
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
// 	user.Password = string(hashedPassword)

// 	// Insert user
// 	_, err = userCollection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		http.Error(w, "Error creating user", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "User registered successfully",
// 	})
// }

// -------------------- LOGIN --------------------
// func Login(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)

// 	filter := bson.M{"email": user.Email}
// 	var foundUser User

// 	err := userCollection.FindOne(context.Background(), filter).Decode(&foundUser)
// 	if err != nil {
// 		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
// 		return
// 	}

// 	// Compare password
// 	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
// 	if err != nil {
// 		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "Login successful",
// 	})
// }
