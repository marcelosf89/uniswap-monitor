package services

import (
	"brahmafi-build-it/monitor/pkg/configurations"
	"brahmafi-build-it/monitor/pkg/models"
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	abiReader abi.ABI
	topics    = [][]common.Hash{
		{crypto.Keccak256Hash([]byte("Swap(address,address,int256,int256,uint160,uint128,int24)"))},
	}
)

type GoEthClient struct {
	client *ethclient.Client
}

func NewGoEthClient() *GoEthClient {
	return &GoEthClient{}
}

func SetupAbiReader() {
	log.Debug().Msg(fmt.Sprintf("abi content file: %v", configurations.Uniswapv3ABI))

	a, err := abi.JSON(strings.NewReader(configurations.Uniswapv3ABI))
	if err != nil {
		log.Fatal().Err(err)
	}

	abiReader = a
}

func (g *GoEthClient) StartClient() {
	rpcEndpoint := configurations.GetRPCEndpoint()
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Fatal().Err(err)
	}

	g.client = client

	SetupAbiReader()
}

func (g *GoEthClient) GetBlockNumber(ctx context.Context) uint64 {
	blockNumber, err := g.client.BlockNumber(ctx)

	if err != nil {
		log.Fatal().Err(err)
	}

	log.Debug().Msg(fmt.Sprintf("Current block number is %v ", blockNumber))

	return blockNumber
}

func (g *GoEthClient) GetLogs(address string, fromBlock uint64, toBlock uint64, ctx context.Context, fn EventSwapFunction) {
	addresses := make([]common.Address, 1)
	addresses[0] = common.HexToAddress(address)

	if fromBlock == toBlock {
		return
	}

	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: addresses,
		Topics:    topics,
	}

	logs, err := g.client.FilterLogs(ctx, query)
	if err != nil {
		log.Debug().Msg(fmt.Sprintf("failed to get logs: %v", err))
		return
	}

	log.Debug().Msg(fmt.Sprintf("--------------------------------------- \n Logs: %v \n --------------------------------------", len(logs)))
	for _, vLog := range logs {

		var event models.SwapEvent
		err := abiReader.UnpackIntoInterface(&event, "Swap", vLog.Data)
		if err != nil {
			log.Debug().Msg(fmt.Sprintf("failed to unpack event data: %v", err))
			continue
		}

		fn(event, vLog.BlockNumber)
	}

}
