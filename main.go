package main

/*

func main() {
	// initializing database and router
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017/default_db?authSource=admin"))

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	// MongoDB collection
	upcomingfightsCollection := client.Database("testing").Collection("fights")
	fightresultsCollection := client.Database("testing").Collection("fights")

	// insert a single document into a collection
	// create a bson.D object
	fights := bson.D{{"time", "fighter 1"}, {"statistics", {}}}
	// insert the bson object using InsertOne()
	result, err := upcomingfightsCollection.InsertOne(context.TODO(), fights)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	fights := []interface{}{
		bson.D{{"time", "fighter 2"}, {"age"}, {"height"}},
		bson.D{{"time", "fighter 3"}, {"age"}, {"height"}},
		bson.D{{"time", "fighter 4"}, {"age"}, {"height"}},
	}
	// insert the bson object slice using InsertMany()
	results, err := upcomingfightsCollection.InsertMany(context.TODO(), fighters)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)

	// server := echo.New()
	// register routes
	// initializing cron jobs
	//      get data
	//      process data and clean up data
	//      store data
}
*/
