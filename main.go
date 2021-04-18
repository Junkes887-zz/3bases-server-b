package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Junkes887/3bases-server-b/controller"
	"github.com/Junkes887/3bases-server-b/database"
	"github.com/Junkes887/3bases-server-b/repository"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

func injects(collection *mongo.Collection, clientRedis *redis.Client, context context.Context) controller.Client {
	time, err := strconv.Atoi(os.Getenv("TIME_MINUTES_REDIS"))

	if err != nil {
		fmt.Println(err)
	}

	rep := repository.Client{
		DB_MONGO:           collection,
		DB_REDIS:           clientRedis,
		CTX:                context,
		TIME_MINUTES_REDIS: time,
	}

	controller := controller.Client{
		DB_MONGO: collection,
		DB_REDIS: clientRedis,
		CTX:      context,
		REP:      rep,
	}

	return controller
}

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	DATABASE := os.Getenv("DATABASE")
	COLLECTION := os.Getenv("COLLECTION")
	context := context.Background()

	dbMongo := database.Context{CTX: context}
	clientMongo := dbMongo.CreateConnectionMongo()
	collection := clientMongo.Database(DATABASE).Collection(COLLECTION)

	clientRedis := database.CreateConnectionRedis()

	controller := injects(collection, clientRedis, context)

	controller.REP.SetDataRedis()

	router := httprouter.New()
	router.GET("/", controller.FindAll)
	router.GET("/:id", controller.Find)
	router.POST("/", controller.Save)
	router.PUT("/:id", controller.Upadate)
	router.DELETE("/:id", controller.Delete)

	c := cors.AllowAll()
	handlerCors := c.Handler(router)

	fmt.Println("Listem " + PORT + ".....")
	http.ListenAndServe(":"+PORT, handlerCors)
}
