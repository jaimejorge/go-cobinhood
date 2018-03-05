package cobinhood

import (
	"fmt"
)

func (c *CobinhoodClient) GetOrderTrades(order_id string) (Trades, error) {
	var Respond Respond

	err := c.Get(fmt.Sprintf("/v1/trading/orders/%s/trades", order_id), &Respond)

	if err != nil {
		return Respond.Result.Trades, err
	}

	return Respond.Result.Trades, nil
}
