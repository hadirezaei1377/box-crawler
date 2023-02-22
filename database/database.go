package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	InsertUpcomingFight(fight map[string]string)
	GetUpcomingFights()
	GetUpcomingFight()
	UpsertFightResults(fight map[string]string)
}

type mongodb struct {
	db *mongo.Client
}

func New(dsn string) (*mongodb, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	return &mongodb{
		db: client,
	}, err
}

func (db mongodb) InsertUpcomingFight(fight map[string]string) {
	db.db.Database("default_db").Collection("upcoming_fights").UpdateOne(context.TODO(), fight, options.Update().SetUpsert(true))
}
func (db mongodb) GetUpcomingFights()                         {}
func (db mongodb) GetUpcomingFight()                          {}
func (db mongodb) UpsertFightResults(fight map[string]string) {}
