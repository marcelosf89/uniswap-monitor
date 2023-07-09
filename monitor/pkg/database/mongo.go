package database

import (
	"context"
	"fmt"
	"uniswap-monitor/monitor/pkg/models"

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

func (db *MongoDb) SaveMonitorState(model models.MonitorState) {
	insertResult, err := db.getCollection(MONITOR_STATE_COLLECTION).InsertOne(db.context, model)
	if err != nil {
		log.Error().Msg(err.Error())
		log.Fatal().Err(err)
	}

	log.Debug().Msg(fmt.Sprintf("Inserted model with ID: %s", insertResult.InsertedID))
}

func (db *MongoDb) SaveDataPoint(model models.DataPoint) {
	insertResult, err := db.getCollection(DATA_POINT_COLLECTION).InsertOne(db.context, model)
	if err != nil {
		log.Error().Msg(err.Error())
		log.Fatal().Err(err)
	}

	log.Debug().Msg(fmt.Sprintf("Inserted model with ID: %s", insertResult.InsertedID))
}

func (db *MongoDb) UpdateMonitorState(id string, blockNumber uint64, token0balance string, token1balance string) {
	update := bson.M{
		"$set": bson.M{
			"lastblock":     blockNumber,
			"token0balance": token0balance,
			"token1balance": token1balance,
		},
	}

	filter := bson.M{"id": id}

	insertResult, err := db.getCollection(MONITOR_STATE_COLLECTION).UpdateOne(db.context, filter, update)
	if err != nil {
		log.Error().Msg(err.Error())
		log.Fatal().Err(err)
	}

	log.Debug().Msg(fmt.Sprintf("Upserted model with ID: %s", insertResult.UpsertedID))
}

func (db *MongoDb) GetMonitorState(field string, key string) *models.MonitorState {

	filter := bson.M{field: key}

	result := &models.MonitorState{}
	err := db.getCollection(MONITOR_STATE_COLLECTION).FindOne(db.context, filter).Decode(&result)
	if err != nil {
		log.Fatal().Err(err)
	}

	return result
}

func (db *MongoDb) GetAllMonitorStates() []models.MonitorState {

	cur, err := db.getCollection(MONITOR_STATE_COLLECTION).Find(db.context, bson.D{})
	if err != nil {
		log.Fatal().Err(err)
	}

	var results []models.MonitorState
	for cur.Next(context.TODO()) {
		var state models.MonitorState
		err := cur.Decode(&state)
		if err != nil {
			log.Fatal().Err(err)
		}

		results = append(results, state)

	}

	return results
}

func (db *MongoDb) getCollection(collection string) *mongo.Collection {
	return db.database.Collection(collection)
}
