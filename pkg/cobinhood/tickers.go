package cobinhood

import (
	"fmt"
)

type Ticker struct {
	TradingPairId  string `json:"Trading_pair_id"`
	Timestamp        int    `json:"Timestamp"`
	H24hhigh        string `json:"24h_high"`
	H24hlow         string `json:"24h_low"`
	H24hopen        string `json:"24h_open"`
	H24hvolume      string `json:"24h_volume"`
	LastTradePrice string `json:"Timestamp"`
	HighestBid      string `json:"Highest_bid"`
	LowestAsk       string `json:"Lowest_ask"`
}

func (c *CobinhoodClient) GetTicker(TradingPairId string) (Ticker, error) {
	var Respond Respond

	err := c.Get(fmt.Sprintf("/v1/market/tickers/%s", TradingPairId), &Respond)

	if err != nil {
		return Respond.Result.Ticker, err
	}

	return Respond.Result.Ticker, nil
}
