package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/go-build/db"
	errorhandler "github.com/swarajkumarsingh/go-build/errorHandler"
	"github.com/swarajkumarsingh/go-build/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(c *gin.Context) {

	defer errorhandler.Recovery(c)

	// Get data off req.body
	var body struct {
		Name  string
		Email string
		Phone string
	}

	c.Bind(&body)

	user := models.User{Name: body.Name, Email: body.Email, Phone: body.Phone}

	// Create a User
	result, err := db.UserDB.InsertOne(context.TODO(), &user)
	errorhandler.HandleError(err, "error while creating user")

	fmt.Println("User", result)

	// Return it in response
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    result,
	})
}

func UpdateUser(c *gin.Context) {

	defer errorhandler.Recovery(c)

	// Get data off req.body
	var body struct {
		Name  string
		Email string
		Phone string
	}

	c.Bind(&body)

	id := c.Param("id")

	user := models.User{Name: body.Name, Email: body.Email, Phone: body.Phone}
	objectId, err := primitive.ObjectIDFromHex(id)
	errorhandler.HandleError(err, "error while creating user")

	// Create a User
	update := bson.M{"$set": bson.M{"Name": user.Name, "Email": user.Email, "Phone": user.Email}}

	result, err := db.UserDB.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
	errorhandler.HandleError(err, "error while creating user")

	if result.ModifiedCount == 0 {
		errorhandler.HandleError(err, "User could not be updated")
		return
	}

	// Return it in response
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func GetUser(c *gin.Context) {
	defer errorhandler.Recovery(c)

	// Get id from params
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	errorhandler.HandleError(err, "error type casting id")

	var foundUser models.User

	// find user based on that id
	err = db.UserDB.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&foundUser)
	errorhandler.HandleError(err, "error while fetching user")


	// Return it in response
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    foundUser,
	})
}

func DeleteUser(c *gin.Context) {
	defer errorhandler.Recovery(c)

	// Get id from params
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	errorhandler.HandleError(err, "error type casting id")

	// find user based on that id
	deleteResult, err := db.UserDB.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	errorhandler.HandleError(err, "error while deleting user")

	if deleteResult.DeletedCount == 0 {
		errorhandler.HandleErrorWithOutError("Could not delete user")
		return
	}

	// Return it in response
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func GetUsers(c *gin.Context) {
	defer errorhandler.Recovery(c)

	var results []models.User
	findOptions := options.Find()
	findOptions.SetLimit(100)

	// find user based on that id
	cursor, err := db.UserDB.Find(context.TODO(), bson.D{{}}, findOptions)
	errorhandler.HandleError(err, "error while fetching user")

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.User
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Return it in response
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"users":   results,
	})
}
