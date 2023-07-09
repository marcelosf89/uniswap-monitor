package services

import (
	"brahmafi-build-it/api/pkg/database"
)

type PoolHistoricItemResponse struct {
	Token0Balance string `json:"token0Balance"`
	Token0Delta   string `json:"token0Delta"`
	Token1Balance string `json:"token1Balance"`
	Token1Delta   string `json:"token1Delta"`
	BlockNumber   uint64 `json:"blockNumber"`
}

func HandleGetPoolHistoric(poolId string) ([]PoolHistoricItemResponse, error) {
	results := database.GetAllDataPointByPoolId(poolId)

	var response []PoolHistoricItemResponse

	for _, v := range results {
		response = append(response, PoolHistoricItemResponse{
			Token0Balance: v.Token0Balance,
			Token0Delta:   v.Token0Balance,
			Token1Balance: v.Token1Balance,
			Token1Delta:   v.Token1Delta,
			BlockNumber:   v.BlockNumber,
		})
	}

	return response, nil
}
