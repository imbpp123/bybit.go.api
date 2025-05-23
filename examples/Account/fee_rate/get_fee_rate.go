package main

import (
	"context"
	"fmt"

	bybit "github.com/imbpp123/bybit.go.api"
)

func main() {
	GetFeeRates()
}

func GetFeeRates() {
	client := bybit.NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", bybit.WithBaseURL(bybit.TESTNET))
	params := map[string]interface{}{"symbol": "BTCUSDT", "category": "linear"}
	accountResult, err := client.NewUtaBybitServiceWithParams(params).GetFeeRates(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}
