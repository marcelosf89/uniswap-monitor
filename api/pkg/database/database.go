package database

import (
	"brahmafi-build-it/api/pkg/models"
	"context"
)

type Database interface {
	Init(connStr string, ctx context.Context)

	OpenConnection()
	CloseConnection()

	Ping() error

	GetAllDataPointByPoolId(poolId string) []models.DataPoint
	GetLatestBlockNumberByPoolId(poolId string) uint64
	GetAllDataPointByPoolIdAndBlocks(poolId string, blocks []uint64) []models.DataPoint
}

var dbImpl Database

func Init(connStr string, db Database, ctx context.Context) {
	dbImpl = db
	dbImpl.Init(connStr, ctx)
}

func OpenConnection() {
	dbImpl.OpenConnection()
}

func CloseConnection() {
	dbImpl.CloseConnection()
}

func Ping() error {
	return dbImpl.Ping()
}

func GetAllDataPointByPoolId(poolId string) []models.DataPoint {
	return dbImpl.GetAllDataPointByPoolId(poolId)
}

func GetLatestBlockNumberByPoolId(poolId string) uint64 {
	return dbImpl.GetLatestBlockNumberByPoolId(poolId)
}

func GetAllDataPointByPoolIdAndBlocks(poolId string, blocks []uint64) []models.DataPoint {
	return dbImpl.GetAllDataPointByPoolIdAndBlocks(poolId, blocks)
}
