package db

import (
	"context"
	"os"

	"github.com/swarajkumarsingh/go-build/colors"
	errorhandler "github.com/swarajkumarsingh/go-build/errorHandler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserDB *mongo.Collection

func ConnectDB() {

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	errorhandler.HandleError(err)

	err = client.Ping(context.TODO(), nil)
	errorhandler.HandleError(err)

	colors.Print(colors.ColorCyan, "Connected to MongoDB!")

	UserDB = client.Database("GoBuild").Collection("User")
}
