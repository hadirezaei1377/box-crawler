package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	InsertUpcomingFight(fight map[string]string) error
	GetUpcomingFights() ([]map[string]string, error)
	GetUpcomingFight(date string)
	UpsertFightResults(fight map[string]string)
}

type mongodb struct {
	db *mongo.Database
}

func New(dsn string) (*mongodb, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	return &mongodb{
		db: client.Database("default_db"),
	}, err
}

func (m mongodb) InsertUpcomingFight(fight map[string]string) error {
	_, err := m.db.Collection("upcoming_fights").UpdateOne(context.TODO(), fight, options.Update().SetUpsert(true))
	return err
}

func (m mongodb) GetUpcomingFights() ([]map[string]string, error) {
	cursor, err := m.db.Collection("upcoming_fights").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	results := make([]map[string]string, 0)
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (m mongodb) GetUpcomingFight(date string) {
	m.db.Collection("upcoming_fights").FindOne(context.TODO(), bson.D{{Key: "date", Value: date}})
}
func (m mongodb) UpsertFightResults(fight map[string]string) {}
