package handlers

import (
	"context"
	"fmt"

	"github.com/dev-gabrielchaves/golang-learning/database"
	"github.com/dev-gabrielchaves/golang-learning/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	cursor, err := database.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		users = append(users, user)
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	var user models.User
	err = database.UserCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return c.Status(404).SendString("User not found")
	}

	return c.JSON(user)
}

// The c *fiber.Ctx is the request context provided by Fiber. It contains request data, headers, and methods for sending responses
func CreateUser(c *fiber.Ctx) error {

	// The user represents an instance of the models.User struct, which represents a user in the system
	var user models.User

	fmt.Printf("BEFORE parsing: %+v \n", user) // Output -> {ID:ObjectID("000000000000000000000000") Name: Email:}

	// Fiber parses the incoming JSON request body and maps it to the user struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Generate a new ObjectID for the user
	user.ID = primitive.NewObjectID()

	fmt.Printf("AFTER parsing: %+v \n", user) // Output -> {ID:ObjectID("60f3b3b3b3b3b3b3b3b3b3b3") Name:Gabriel Email:gabriel@gmail.com

	// Insert the user into the MongoDB collection
	_, err := database.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Return the created user as a JSON response
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	update := bson.M{"$set": user}
	_, err = database.UserCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("User updated successfully")
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	_, err = database.UserCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("User deleted successfully")
}
