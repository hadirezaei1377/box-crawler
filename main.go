package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// initializing database and router
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017/default_db?authSource=admin"))
	if err != nil {
		panic(err)
	}

	// server := echo.New()
	// register routes
	// initializing cron jobs
	//      get data
	//      process data and clean up data
	//      store data
}
