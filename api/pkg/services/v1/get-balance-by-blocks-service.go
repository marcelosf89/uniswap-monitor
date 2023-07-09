package services

import (
	"brahmafi-build-it/api/pkg/database"
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

const (
	BLOCK_LATEST = "latest"
)

type PoolBalanceItemResponse struct {
	Token0Balance string `json:"token0Balance"`
	Token1Balance string `json:"token1Balance"`
	Tick          uint64 `json:"tick"`
	BlockNumber   uint64 `json:"blockNumber"`
}

func HandleGetPoolBalanceByBlocks(poolId string, blocks []string) ([]PoolBalanceItemResponse, error) {

	blocksSanitized := getAllBlockNumbers(poolId, blocks)
	results := database.GetAllDataPointByPoolIdAndBlocks(poolId, blocksSanitized)

	var response []PoolBalanceItemResponse

	for _, v := range results {
		response = append(response, PoolBalanceItemResponse{
			Token0Balance: v.Token0Balance,
			Token1Balance: v.Token1Balance,
			Tick:          v.Tick,
			BlockNumber:   v.BlockNumber,
		})
	}

	return response, nil
}

func getAllBlockNumbers(poolId string, blocks []string) []uint64 {
	var result []uint64

	for _, v := range blocks {
		if v == BLOCK_LATEST {
			latestBlock := database.GetLatestBlockNumberByPoolId(poolId)
			result = append(result, latestBlock)
		} else {
			num, err := strconv.ParseUint(v, 10, 64)

			if err != nil {
				log.Debug().Msg(fmt.Sprintf("Error on convert block string to uint, value %v", v))
				continue
			}

			result = append(result, num)
		}
	}

	return result
}
