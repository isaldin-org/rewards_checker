package main

import (
	"fmt"
	"math/big"
	"os"
)

var (
	infuraConnectionString = fmt.Sprintf("https://mainnet.infura.io/v3/%s", os.Getenv("INFURA_API_KEY"))
	contractAddress        = os.Getenv("ADDRESS")
	oneInchTokenAddress    = "0x111111111117dc0aa78b770fa6a738034120c302"
	userWallet             = os.Getenv("WALLET")
	botToken               = os.Getenv("BOT_TOKEN")
	allowedUserId          = os.Getenv("ALLOWED_USER_ID")
)

func main() {
	reward := GetReward()

	SendToTelegram(weiToUnit(reward))
}

func weiToUnit(weis big.Int) string {
	weisFloat, err := new(big.Float).SetString(weis.String())
	if !err {
		panic(err)
	}
	units := new(big.Float).Quo(weisFloat, new(big.Float).SetFloat64(1e18))
	return units.Text('f', 2)
}
