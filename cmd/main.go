package main

import (
	"fmt"
	resty "github.com/go-resty/resty/v2"
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
	coinApiKey             = os.Getenv("COINAPI_KEY")
)

type Rate struct {
	Time           string
	Asset_id_base  string
	Asset_id_quote string
	Rate           float64
}

func main() {
	rate, err := getOneInchTokenPrice()
	if err != nil {
		// so dirty retry :)
		rate, err = getOneInchTokenPrice()
		if err != nil {
			rate = *new(big.Float).SetUint64(0)
		}
	}

	reward := GetReward()
	usdtPrice := getUsdPrice(weiToUnit(reward), rate)
	usdtPriceString := usdtPrice.Text('f', 2)

	SendToTelegram(weiToUnitString(reward), usdtPriceString)
}

func weiToUnitString(weis big.Int) string {
	units := weiToUnit(weis)
	return units.Text('f', 2)
}

func weiToUnit(weis big.Int) big.Float {
	weisFloat, success := new(big.Float).SetString(weis.String())
	if !success {
		panic(!success)
	}
	units := new(big.Float).Quo(weisFloat, new(big.Float).SetFloat64(1e18))
	return *units
}

func getOneInchTokenPrice() (big.Float, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("X-CoinAPI-Key", coinApiKey).
		SetResult(Rate{}).
		ForceContentType("application/json").
		Get("https://rest.coinapi.io/v1/exchangerate/1INCH/USDT")

	if err != nil {
		return *new(big.Float).SetInt64(0), err
	}

	rate, ok := resp.Result().(*Rate)

	if !ok {
		return *new(big.Float).SetInt64(0), err
	}

	return *new(big.Float).SetFloat64(rate.Rate), nil
}

func getUsdPrice(rewards big.Float, rate big.Float) big.Float {
	return *new(big.Float).Mul(&rate, &rewards)
}
