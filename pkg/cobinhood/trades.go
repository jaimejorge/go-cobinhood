package cobinhood

import (
	"fmt"
)

type Trade struct {
	Id         string `json:"id"`
	Price      string `json:"price"`
	Size       string `json:"size"`
	MakerSide string `json:"maker_side"`
	Timestamp  jsonTime `json:"timestamp"`
}
type Trades []Trade

func (c *CobinhoodClient) GetTrades(TradingPairId string) (Trades, error) {
	var Respond Respond

	err := c.Get(fmt.Sprintf("/v1/market/trades/%s", TradingPairId), &Respond)

	if err != nil {
		return Respond.Result.Trades, err
	}

	return Respond.Result.Trades, nil
}
