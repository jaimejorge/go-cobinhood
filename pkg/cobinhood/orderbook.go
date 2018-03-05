package cobinhood

import "fmt"

type Orderbook struct {
	Sequence int  `json:"sequence"`
	Bids     [][]string   `json:"bids"`
	Asks     [][]string  `json:"asks"`
}

func (c *CobinhoodClient) GetOrderbooks(trading_pair_id string) (Orderbook, error) {
	var Respond Respond

	err := c.Get(fmt.Sprintf("/v1/market/orderbooks/%s", trading_pair_id), &Respond)

	if err != nil {
		return Respond.Result.Orderbook, err
	}

	return Respond.Result.Orderbook, nil
}
