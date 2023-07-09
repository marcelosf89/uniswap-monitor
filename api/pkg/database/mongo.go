package database

import (
	"brahmafi-build-it/api/pkg/models"
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE_NAME = "monitor"

	MONITOR_STATE_COLLECTION = "MonitorState"
	DATA_POINT_COLLECTION    = "DataPoint"
)

type MongoDb struct {
	ConnString string
	client     *mongo.Client
	database   *mongo.Database
	context    context.Context
}

func NewMongoDb() *MongoDb {
	return &MongoDb{}
}

func (db *MongoDb) Init(connStr string, ctx context.Context) {
	db.ConnString = connStr
	db.context = ctx
}

func (db *MongoDb) OpenConnection() {
	clientOptions := options.Client().ApplyURI(db.ConnString)
	client, err := mongo.Connect(db.context, clientOptions)
	if err != nil {
		log.Fatal().Err(err)
	}

	err = client.Ping(db.context, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("Connected to MongoDB!")

	db.client = client
	db.database = client.Database(DATABASE_NAME)
}

func (db *MongoDb) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	return db.client.Ping(ctx, nil)
}

func (db *MongoDb) CloseConnection() {
	if db.client.Ping(db.context, nil) == nil {
		err := db.client.Disconnect(db.context)
		if err != nil {
			log.Error().Err(err).Msgf("error on trying to close MongoDb connection: %v", err)
			return
		}
		log.Info().Msg("Database connection has been closed.")
	}
}

func (db *MongoDb) GetAllDataPointByPoolId(poolId string) []models.DataPoint {

	filter := bson.M{"poolid": poolId}
	cur, err := db.getCollection(DATA_POINT_COLLECTION).Find(db.context, filter)
	if err != nil {
		log.Fatal().Err(err)
	}

	var results []models.DataPoint
	for cur.Next(context.TODO()) {
		var dataPoint models.DataPoint
		err := cur.Decode(&dataPoint)
		if err != nil {
			log.Fatal().Err(err)
		}

		results = append(results, dataPoint)

	}

	return results
}

func (db *MongoDb) GetLatestBlockNumberByPoolId(poolId string) uint64 {

	filter := bson.M{"address": poolId}

	result := &models.MonitorState{}
	err := db.getCollection(MONITOR_STATE_COLLECTION).FindOne(db.context, filter).Decode(&result)
	if err != nil {
		log.Fatal().Err(err)
	}

	return result.LastBlock
}

func (db *MongoDb) GetAllDataPointByPoolIdAndBlocks(poolId string, blocks []uint64) []models.DataPoint {

	filter := bson.M{
		"poolid":      poolId,
		"blocknumber": bson.M{"$in": blocks},
	}

	cur, err := db.getCollection(DATA_POINT_COLLECTION).Find(db.context, filter)
	if err != nil {
		log.Fatal().Err(err)
	}

	var results []models.DataPoint
	for cur.Next(context.TODO()) {
		var dataPoint models.DataPoint
		err := cur.Decode(&dataPoint)
		if err != nil {
			log.Fatal().Err(err)
		}

		results = append(results, dataPoint)
	}

	return results
}

func (db *MongoDb) getCollection(collection string) *mongo.Collection {
	return db.database.Collection(collection)
}
