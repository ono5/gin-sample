// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin.
//
//	Schemes: http
//  Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//	Contact: Mohamed Labouardy <mohamed@labouardy.com> https://labouardy.com
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ono5/gin-sample/recipes-api/handlers"
	"github.com/ono5/gin-sample/recipes-api/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler
var authHandler *handlers.AuthHandler

func init() {
	ctx := context.Background()
	log.Println(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	status := redisClient.Ping()
	fmt.Println("Connection To Redis: ", status)

	recipesHandler = handlers.NewRecipeHandler(ctx, collection, redisClient)

	collectionUsers := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
}

func main() {
	router := gin.Default()
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/refresh", authHandler.RefreshHandler)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
	authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	authorized.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	// router.DELETE("/recipes/:id", DeleteRecipeHandler)
	// router.GET("/recipes/:id", GetRecipeHandler)
	router.Run()
}
