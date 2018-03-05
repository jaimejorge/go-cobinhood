package cobinhood

import (
	"fmt"
	"strings"
)

type Order struct {
	Id           string   `json:"id"`
	TradingPair string `json:"trading_pair"`
	State        string `json:"state"`
	Side         string `json:"side"`
	Type         string `json:"type"`
	Price        string  `json:"price"`
	Size         string  `json:"size"`
	Filled       string  `json:"filled"`
	Timestamp    jsonTime `json:"timestamp"`
}

func (c *CobinhoodClient) GetOrder(OrderId string) (Order, error) {
	var Respond Respond

	err := c.Get(fmt.Sprintf("/v1/trading/orders/%s", OrderId), &Respond)

	if err != nil {
		return Respond.Result.Order, err
	}

	return Respond.Result.Order, nil
}

func (c *CobinhoodClient) GetAllOrders() ([]Order, error) {
	var Respond Respond

	err := c.Get("/v1/trading/orders", &Respond)

	if err != nil {
		return Respond.Result.Orders, err
	}

	return Respond.Result.Orders, nil
}

func (c *CobinhoodClient) GetHistoryOrders(trading_pair_ids []string, limit string) ([]Order, error) {
	if limit == "" {
		limit = "50"
	}
	var Respond Respond
	path := ""
	if len(trading_pair_ids) == 0 {
		path = "/v1/trading/order_history"
	} else {
		path = fmt.Sprintf("/v1/trading/order_history?limit=%s&trading_pair_ids=%s", limit, strings.Join(trading_pair_ids[:], "&trading_pair_ids="))
	}
	err := c.Get(path, &Respond)

	if err != nil {
		return Respond.Result.Orders, err
	}

	return Respond.Result.Orders, nil
}
