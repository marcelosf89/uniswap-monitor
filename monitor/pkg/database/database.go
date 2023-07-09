package database

import (
	"context"
	"uniswap-monitor/monitor/pkg/models"
)

type Database interface {
	Init(connStr string, ctx context.Context)

	OpenConnection()
	CloseConnection()

	SaveDataPoint(model models.DataPoint)

	SaveMonitorState(model models.MonitorState)
	UpdateMonitorState(id string, blockNumber uint64, token0balance string, token1balance string)
	GetMonitorState(field string, key string) *models.MonitorState
	GetAllMonitorStates() []models.MonitorState
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

func SaveMonitorState(model models.MonitorState) {
	dbImpl.SaveMonitorState(model)
}

func SaveDataPoint(model models.DataPoint) {
	dbImpl.SaveDataPoint(model)
}

func UpdateMonitorState(id string, blockNumber uint64, token0balance string, token1balance string) {
	dbImpl.UpdateMonitorState(id, blockNumber, token0balance, token1balance)
}

func GetMonitorState(field string, key string) *models.MonitorState {
	return dbImpl.GetMonitorState(field, key)
}

func GetAllMonitorStates() []models.MonitorState {
	return dbImpl.GetAllMonitorStates()
}
