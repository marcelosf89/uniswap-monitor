package services

import (
	"brahmafi-build-it/monitor/pkg/models"
	"context"
	"time"
)

type Monitor interface {
	StartMonitor(state *models.MonitorState, ctx context.Context)
}

type EthClient interface {
	StartClient()
	GetBlockNumber(ctx context.Context) uint64
	GetLogs(address string, fromBlock uint64, toBlock uint64, ctx context.Context, fn EventSwapFunction)
}

type EventSwapFunction func(models.SwapEvent, uint64)

var ethClientImpl EthClient

func InitEthClient(client EthClient) {
	ethClientImpl = client
}

func StartMonitor(monitor Monitor, state *models.MonitorState, ctx context.Context) {
	monitor.StartMonitor(state, ctx)
}

func StartEthClient() {
	ethClientImpl.StartClient()
}

func GetBlockNumber() uint64 {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return ethClientImpl.GetBlockNumber(ctx)
}

func GetLogs(address string, fromBlock uint64, toBlock uint64, fn EventSwapFunction) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ethClientImpl.GetLogs(address, fromBlock, toBlock, ctx, fn)
}
