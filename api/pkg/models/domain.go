package models

type Model interface {
	GetId() string
}

type MonitorState struct {
	Id            string
	Address       string
	LastBlock     uint64
	Token0Balance string
	Token1Balance string
	Enabled       bool
}

func (model *MonitorState) GetId() string {
	return model.Id
}

type DataPoint struct {
	Id            string
	PoolId        string
	BlockNumber   uint64
	Tick          uint64
	Token0Balance string
	Token0Delta   string
	Token1Balance string
	Token1Delta   string
}

func (model *DataPoint) GetId() string {
	return model.Id
}
