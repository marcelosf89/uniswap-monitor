package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"uniswap-monitor/monitor/pkg/models"
)

type MockDatabase struct {
	InitCalled                bool
	OpenConnectionCalled      bool
	CloseConnectionCalled     bool
	SaveCalled                bool
	UpdateMonitorStateCalled  bool
	GetMonitorStateCalled     bool
	GetAllMonitorStatesCalled bool
	ExpectedDataPoint         models.DataPoint
	ExpectedMonitorState      models.MonitorState
	ExpectedField             string
	ExpectedKey               string
}

func (db *MockDatabase) Init(connStr string, ctx context.Context) {
	db.InitCalled = true
}

func (db *MockDatabase) OpenConnection() {
	db.OpenConnectionCalled = true
}

func (db *MockDatabase) CloseConnection() {
	db.CloseConnectionCalled = true
}

func (db *MockDatabase) SaveDataPoint(model models.DataPoint) {
	db.SaveCalled = true
	db.ExpectedDataPoint = model
}

func (db *MockDatabase) SaveMonitorState(model models.MonitorState) {
	db.SaveCalled = true
	db.ExpectedMonitorState = model
}

func (db *MockDatabase) UpdateMonitorState(id string, blockNumber uint64, token0balance string, token1balance string) {
	db.UpdateMonitorStateCalled = true
}

func (db *MockDatabase) GetMonitorState(field string, key string) *models.MonitorState {
	db.GetMonitorStateCalled = true
	db.ExpectedField = field
	db.ExpectedKey = key
	return &models.MonitorState{}
}

func (db *MockDatabase) GetAllMonitorStates() []models.MonitorState {
	db.GetAllMonitorStatesCalled = true
	return []models.MonitorState{}
}

func TestSaveMonitorState(t *testing.T) {
	// Arrange
	mockDB := &MockDatabase{}

	Init("mongodb://localhost:27017", mockDB, context.Background())

	testModel := models.MonitorState{
		Id: "1",
	}

	// Act
	SaveMonitorState(testModel)

	// Assert
	assert.True(t, mockDB.SaveCalled)
	assert.Equal(t, testModel, mockDB.ExpectedMonitorState)
}

func TestSaveDataPoint(t *testing.T) {
	// Arrange
	mockDB := &MockDatabase{}

	Init("mongodb://localhost:27017", mockDB, context.Background())

	testModel := models.DataPoint{
		Id: "1",
	}

	// Act
	SaveDataPoint(testModel)

	// Assert
	assert.True(t, mockDB.SaveCalled)
	assert.Equal(t, testModel, mockDB.ExpectedDataPoint)
}

func TestUpdateMonitorState(t *testing.T) {
	// Arrange
	mockDB := &MockDatabase{}

	Init("mongodb://localhost:27017", mockDB, context.Background())

	// Act
	UpdateMonitorState("1", 123456, "100", "200")

	// Assert
	assert.True(t, mockDB.UpdateMonitorStateCalled)
}

func TestGetMonitorState(t *testing.T) {
	// Arrange
	mockDB := &MockDatabase{}

	Init("mongodb://localhost:27017", mockDB, context.Background())

	// Act
	GetMonitorState("field", "key")

	// Assert
	assert.True(t, mockDB.GetMonitorStateCalled)
	assert.Equal(t, "field", mockDB.ExpectedField)
	assert.Equal(t, "key", mockDB.ExpectedKey)
}

func TestGetAllMonitorStates(t *testing.T) {
	// Arrange
	mockDB := &MockDatabase{}

	Init("mongodb://localhost:27017", mockDB, context.Background())

	// Act
	GetAllMonitorStates()

	// Assert
	assert.True(t, mockDB.GetAllMonitorStatesCalled)
}
