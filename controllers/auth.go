package controllers

import (
	"context"
	"log"

	"Go_auth_app/database"
	"Go_auth_app/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Register a new user(handle function)
func Register(c *fiber.Ctx) error {
	var user models.User

	// Parsing the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validating fields
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
	}

	// Checking if the user already exists
	var existingUser models.User
	err := database.UsersCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		// if exists return
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email is already registered"})
	}

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	// Inserting the user into the database
	_, err = database.UsersCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error inserting user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Registration successful"})
}

// Login a user(handle function)
func Login(c *fiber.Ctx) error {
	var input models.User

	// Parsing the request body
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validating fields
	if input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email and password are required"})
	}

	// Finding the user by email
	var user models.User
	err := database.UsersCollection.FindOne(context.Background(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User Not Registererd"})
	}

	// Checking the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		// if password do not match return 
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong Password"})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful"})
}
