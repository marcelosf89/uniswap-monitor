package services

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"uniswap-monitor/monitor/pkg/database"
	"uniswap-monitor/monitor/pkg/models"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const (
	WAITING_BLOCKS_NUMBERS = 12
)

type UniswapV3Monitor struct {
	state *models.MonitorState
}

func NewUniswapV3Monitor() *UniswapV3Monitor {
	return &UniswapV3Monitor{}
}

func (monitor *UniswapV3Monitor) StartMonitor(state *models.MonitorState, ctx context.Context) {
	monitor.state = state

	if !state.Enabled {
		log.Info().Msg(fmt.Sprintf("Skip monitoring for address '%s' ...", state.Address))
		return
	}

	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {

			monitor.getLog()
			time.Sleep(time.Second * 10)

			select {
			case <-ctx.Done():
				log.Info().Msg(fmt.Sprintf("Monitoring stopping for address '%s' ...", state.Address))
				return
			default:
			}
		}
	}()
}

func (monitor *UniswapV3Monitor) addDataPoint(current models.SwapEvent, blockNumber uint64) {

	token0, _ := new(big.Int).SetString(monitor.state.Token0Balance, 10)
	token1, _ := new(big.Int).SetString(monitor.state.Token1Balance, 10)

	database.SaveDataPoint(models.DataPoint{
		Id:            fmt.Sprint(uuid.New()),
		PoolId:        monitor.state.Address,
		BlockNumber:   blockNumber,
		Tick:          current.Tick.Uint64(),
		Token0Balance: current.Amount0.String(),
		Token0Delta:   current.Amount0.Add(current.Amount0, token0).String(),
		Token1Balance: current.Amount1.String(),
		Token1Delta:   current.Amount0.Add(current.Amount0, token1).String(),
	})
}

func (monitor *UniswapV3Monitor) updateState(token0Balance *big.Int, token1Balance *big.Int, blockNumber uint64) {
	token0, _ := new(big.Int).SetString(monitor.state.Token0Balance, 10)
	token1, _ := new(big.Int).SetString(monitor.state.Token1Balance, 10)

	monitor.state.Token0Balance = token0Balance.Add(token0Balance, token0).String()
	monitor.state.Token1Balance = token1Balance.Add(token1Balance, token1).String()

	if monitor.state.LastBlock < blockNumber {
		monitor.state.LastBlock = blockNumber
	}

	database.UpdateMonitorState(monitor.state.Id, monitor.state.LastBlock, monitor.state.Token0Balance, monitor.state.Token1Balance)
}

func (monitor *UniswapV3Monitor) getLog() {
	blockNumber := GetBlockNumber()

	if monitor.state.LastBlock+WAITING_BLOCKS_NUMBERS > blockNumber {
		return
	}

	log.Debug().Msg(fmt.Sprintf("Getting blocks from %v to %v", monitor.state.LastBlock, blockNumber))
	GetLogs(monitor.state.Address, monitor.state.LastBlock, blockNumber, func(current models.SwapEvent, blockNumber uint64) {
		monitor.updateState(current.Amount0, current.Amount1, blockNumber)
		monitor.addDataPoint(current, blockNumber)
		log.Debug().Msg(fmt.Sprintf("Block Number %v updated", blockNumber))
	})

	monitor.state.LastBlock = blockNumber
}
