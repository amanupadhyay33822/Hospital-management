package handlers

import (
	"encoding/json"
	"go-auth/internal/auth"
	"go-auth/internal/config"
	"go-auth/internal/models"
	"go-auth/internal/utils"
	"go-auth/internal/validation"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var input validation.Auth

	var err error
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.SendResponse(w, false, "Invalid JSON", nil)
		return
	}

	err = validation.Validateregister(input)
	if err != nil {
		utils.SendResponse(w, false, validation.FormatValidationError(err), nil)
		return
	}

	var collection *mongo.Collection
	collection = config.GetCollection(config.Client, "users")

	var filter bson.M
	filter = bson.M{"email": input.Email}
	var existingUser models.User

	ctx := r.Context()
	err = collection.FindOne(ctx, filter).Decode(&existingUser)
	if err == nil {
		utils.SendResponse(w, false, "User already exists", nil)
		return
	}
	if err != mongo.ErrNoDocuments {
		utils.SendResponse(w, false, "Database error", nil)
		return
	}

	var hashedPassword string
	hashedPassword, err = utils.HashPassword(input.Password)
	if err != nil {
		utils.SendResponse(w, false, "Failed to hash password", nil)
		return
	}

	var newUser models.User
	newUser = models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		utils.SendResponse(w, false, "Failed to create user", nil)
		return
	}

	utils.SendEmailAsync(utils.EmailData{
		To:      input.Email,
		Subject: "Welcome to Our Platform",
		Body:    "Thank you for registering!",
	})

	utils.SendResponse(w, true, "User registered successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input validation.Login

	// decode json
	var err error
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.SendResponse(w, false, "Invalid JSON", nil)
		return
	}

	// validate input
	err = validation.ValidateLogin(input)
	if err != nil {
		utils.SendResponse(w, false, validation.FormatValidationError(err), nil)
		return
	}

	//get user collection
	var collection *mongo.Collection
	collection = config.GetCollection(config.Client, "users")
	var user models.User
	//find user by email

	ctx := r.Context()
	err = collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		utils.SendResponse(w, false, "Invalid email or password", nil)
		return
	}
	if err != nil {
		utils.SendResponse(w, false, "Database error", nil)
		return
	}
	// compare password (brcypt)
	err = utils.CheckPassword(user.Password, input.Password)
	if err != nil {
		utils.SendResponse(w, false, "Invalid email or password", nil)
		return
	}
	//Generate JWT (optionl)
	var token string
	token, err = auth.GenerateToken(user.Email, user.Role)
	if err != nil {
		utils.SendResponse(w, false, "Failed to generate token", nil)
		return
	}
	// Return sucess response
	utils.SendResponse(w, true, "Login successful", token)
}
