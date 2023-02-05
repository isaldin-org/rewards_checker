package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetReward() big.Int {
	client, err := ethclient.Dial(infuraConnectionString)
	if err != nil {
		panic("Couldn't connect to Infura")
	}

	address := common.HexToAddress(contractAddress)
	instance, err := NewRewards(address, client)
	if err != nil {
		panic("Couldn't initialize contract")
	}

	farmed, _ := instance.Farmed(nil, common.HexToAddress(oneInchTokenAddress), common.HexToAddress(userWallet))
	return *farmed
}
